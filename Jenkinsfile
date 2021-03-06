// HELPER FUNCTIONS
def commitSha1() {
  sh 'git rev-parse HEAD > commitSha1'
  def commit = readFile('commitSha1').trim()
  sh 'rm commitSha1'
  // Use first 11 chars from full SHA1
  commit.substring(0, 11)
}

// STAGE FUNCTIONS DEFINITION
def checkoutCode(Map args=[:]){
  stage('Clone repository') {
    // CHECKOUT CODE REPO
    if(args.checkout_tag){
      scmVars = checkout([
          $class: 'GitSCM',
          branches: [[name: "refs/tags/${SELECTED_TAG}"]],
          doGenerateSubmoduleConfigurations: false,
          extensions: [[$class: 'CleanCheckout'],
          [$class: 'PruneStaleBranch']],
          submoduleCfg: []
      ])
      // Re-init variables
      prepareVars()
    } else {
      scmVars = checkout scm
    }
  }
}

def getEnvName() {
  def envName = ""
  if(env.TAG_NAME){
    return "tag"
  }
  switch(env.BRANCH_NAME) {
    case "develop":
      return "dev"
      break
    case "sit":
      return "sit"
      break
    case "uat":
      return "uat"
      break
    default:
      if(env.SELECTED_TAG != null && SELECTED_TAG != "") {
        return "prod"
      } else {
        throw new Exception("Cannot define environment name! (raw: ${env.BRANCH_NAME})")
      }
  }
}

def prepareVars() {
  appName = 'meeting-room-service'
  sonarqubeName ="touch-meeting-room-${appName}"
  commitHash = commitSha1()
  envName = getEnvName()
  appFullName = "${appName}-${envName}"
  echo "This job running on [${envName}] configurations."
  // Docker registry
  imgRepoServerUrl = "registry.touchdevops.com"
  imgRepoCred = "nexus-userpass-jenkins"
  imgRepoName = "product-meeting-room"
  imgFullName = "${imgRepoServerUrl}/${imgRepoName}/${appName}"
  imgTag = "${envName}-${commitHash}"
  // HELM CHARTS
  helmRepoUrl = "https://git.touchdevops.com/infrastructure/helm-charts.git"
  helmRepoCred = "gitlab-userpass-jenkins"
  helmSubDir = "helm-charts"
  helmChartName = "golang"
  helmValuesFile = "${WORKSPACE}/.helmValues/${envName}.yaml"
  helmWaitTimeout = "5m"
  // Tests variables
  switch(envName) {
    case "dev":
      testBaseUrl = "https://api.touchdevops.com/dev/touchgo" // Copied from touch.go, should replace.
      robotTestFile = "sample.robot"
      bztTestFile = "tests/Performance/taurus.yml"
      break
    case "sit":
      testBaseUrl = "https://api.touchdevops.com/sit/touchgo" // Copied from touch.go, should replace.
      bztTestFile = "tests/Performance/taurus.yml"
      break
    case "uat":
      testBaseUrl = "https://api.touchdevops.com/uat/touchgo" // Copied from touch.go, should replace.
      bztTestFile = "tests/Performance/taurus.yml"
      break
    case "tag":
      imgTag = env.TAG_NAME
      break
    case "prod":
      imgTag = SELECTED_TAG // FROM JENKINS JOB PARAMETERS
      break
  }
}

def runUnitTest() {
  stage('Unit Test') {
    container('golang'){
      env.APP_ENV = "dev"
      env.MONGODB_ENDPOINT = "mongodb://root:UZiYgNT6ZeBo9Jxs@mongodb-dev:27017"
      env.MONGODB_NAME = "authentication_unit"
      env.MONGODB_USERS_TABLE_NAME = "users"
      env.APP_NAME =  "authentication-service"
      env.TIMEZONE =  "asia/bangkok"
      env.JAEGER_AGENT_HOST =  "localhost"
      env.JAEGER_AGENT_PORT =  "6831"
      env.UID_HEADER_NAME =  "X-Authenticated-Userid"
      env.SWAGGER_HOST =  "mediator.touchdevops.com/cms/content/dev"
      env.BASE_PATHS =  "/api/v1"
      env.MESSAGE_BROKER_BACKOFF_TIME = "2"
      env.MESSAGE_BROKER_MAXIMUM_RETRY = "3"
      env.MESSAGE_BROKER_ENDPOINT = "kafka-infra.infra:9092"
      env.MESSAGE_BROKER_GROUP = "authentication"
      env.MESSAGE_BROKER_VERSION = "7.0.3"

      sh "env"
      sh "go mod tidy"
      sh "go test -coverprofile=coverage.out ./..."
      sh "go tool cover -func=coverage.out"
    }
  }
}

def runOWASP() {
  stage('OWASP Scanner') {
    container('jnlp'){
      dir ("tests"){
        dependencycheck(
          additionalArguments: "--out dependency-check-report.xml",
          odcInstallation: "owasp-scanner"
        )
        dependencyCheckPublisher(
          pattern: 'dependency-check-report.xml'
        )
      }
    }
  }
}

def runSonarQube() {
  stage('SonarQube Analysis') {
    container('jnlp'){
      def scannerHome = tool 'sonarqube-scanner';
      withSonarQubeEnv('sonarqube') {
        // SCAN AND SUBMIT COVERAGE RESULT FROM JEST
        def sonarOptions = []
        sonarOptions.add("-Dsonar.projectKey=${sonarqubeName}") // SET PROJECT KEY
        sonarOptions.add("-Dsonar.projectName=${sonarqubeName}") // SET PROJECT NAME
        sonarOptions.add("-Dsonar.sources=.") // SET SOURCE PATH
        sonarOptions.add("-Dsonar.exclusions=**/mocks/**,**/test/**,setup.go,main.go,docs/docs.go,**/repository/**,**/domain**,**/config/**,Jenkinsfile,**/.helmValues/**,**/.github/**,**/deployment/**,**/development/**/,README.md,SECURITY.md,.gitignore,Dockerfile") // EXCLUDE TESTS PATH
        sonarOptions.add("-Dsonar.tests=.")
        sonarOptions.add("-Dsonar.test.inclusions=*/_test.go")
        sonarOptions.add("-Dsonar.go.coverage.reportPaths=coverage.out")
//         sonarOptions.add("-Dsonar.junit.reportPaths=tests/junit.xml") // READ JUNIT REPORTx
        sonarOptions = sonarOptions.join(' ')
        sh "${scannerHome}/bin/sonar-scanner ${sonarOptions}"
      }
    }
  }
  stage('Quality Gate') {
    container('jnlp'){
      // WAITING FOR SONARQUBE QUALITY GATE RESULT
      timeout(time: 1, unit: 'HOURS') { // Just in case something goes wrong, pipeline will be killed after a timeout
        def qg = waitForQualityGate() // Reuse taskId previously collected by withSonarQubeEnv
        if (qg.status != 'OK') {
          error "Pipeline aborted due to quality gate failure: ${qg.status}"
        }
      }
    }
  }
}

def buildAndPushDockerImage() {
  stage('Build Image') {
    container('docker'){
      echo "Start building image [${imgFullName}:${imgTag}]"
      docker.withRegistry("https://${imgRepoServerUrl}", "${imgRepoCred}") {
        def img = docker.build("${imgFullName}:${imgTag}", "--pull .")
        img.push()
        if(envName != "tag"){
          img.push("${envName}-latest")
        }
      }
    }
  }
}

def deployApp() {
  stage('Deploy') {
    echo "Deploy"
    container("helm") {
      checkout changelog: false, poll: false, scm: [
        $class: 'GitSCM',
        branches: [[name: '*/master']],
        doGenerateSubmoduleConfigurations: false,
        extensions: [
          [$class: 'CleanCheckout'],
          [$class: 'CloneOption', depth: 0, noTags: true, reference: '', shallow: true],
          [$class: 'SubmoduleOption', disableSubmodules: true, parentCredentials: false, recursiveSubmodules: false, reference: '', trackingSubmodules: false],
          [$class: 'RelativeTargetDirectory', relativeTargetDir: helmSubDir]],
        submoduleCfg: [],
        userRemoteConfigs: [[credentialsId: helmRepoCred, url: helmRepoUrl]]
      ]
      dir(helmSubDir) {
        sh "helm version"
        def helmOptions = "--namespace ${envName} --set-string \"image.repository=${imgFullName},image.tag=${imgTag}\""
        sh "helm upgrade -i ${appFullName} -f ${helmValuesFile} ${helmOptions} --wait --timeout ${helmWaitTimeout} ${helmChartName}"
      }
    }
  }
}

def runPerfTest() {
  stage('Performance Test (JMeter)') {
    container('bzt'){
      // Enable this if you have a test config.
      // bzt params: "-o settings.env.BASE_URL=\"${testBaseUrl}\" ${bztTestFile}", alwaysUseVirtualenv: false
    }
  }
}

def slaveLabel = 'golang1-13'
def slaveCloudName = 'touch-nonprd'

podTemplate(
  label: slaveLabel,
  cloud: slaveCloudName
){
  node(slaveLabel) {

    checkoutCode()
    prepareVars()

    switch(envName) {
      case ["dev", "sit", "uat"]:
//         runUnitTest()
//         runOWASP()
//         runSonarQube()
        buildAndPushDockerImage()
        deployApp()
        // runPerfTest()
        break
      case "tag":
        buildAndPushDockerImage()
        break
      case "prod":
        checkoutCode(checkout_tag: true)
        deployApp()
        break
    }
  }
}
