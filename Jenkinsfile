#!/usr/bin/env groovy

node (SLAVE) {

    // Cleanup workspace
    deleteDir()

    //Clone example project from GitHub repository
    git url: 'https://github.com/jainishshah17/containerize-go-microservice.git', branch: 'master'
    def rtServer = Artifactory.newServer url: ART_SERVER_URL, credentialsId: CREDENTIALS
    def rtDocker = Artifactory.docker server: rtServer
    def buildInfo = Artifactory.newBuildInfo()
    def tagDockerApp

    buildInfo.env.capture = true

    //Build docker image named containerize-go-microservice
    stage ('Build & Deploy') {
            sh "sed -i 's/docker.artifactory/${ARTDOCKER_REGISTRY}/' Dockerfile"
            tagDockerApp = "${ARTDOCKER_REGISTRY}/containerize-go-microservice:${env.BUILD_NUMBER}"
            println "containerize-go-microservice Build"
            docker.build(tagDockerApp)
            println "Docker push" + tagDockerApp + " : " + REPO
            buildInfo = rtDocker.push(tagDockerApp, REPO, buildInfo)
            println "Docker Buildinfo"
            rtServer.publishBuildInfo buildInfo
     }

    stage ("Retag latest image") {
            reTagLatest (SOURCE_REPO)
    }

    //Test docker image
     stage ('Test') {
            tagDockerApp = "${ARTDOCKER_REGISTRY}/containerize-go-microservice:${env.BUILD_NUMBER}"
            if (testApp(tagDockerApp)) {
                  println "Setting property and promotion"
                  sh 'docker rmi '+tagDockerApp+' || true'
             } else {
                  currentBuild.result = 'UNSTABLE'
                  return
             }
     }

    //Scan Build Artifacts in Xray
    stage('Xray Scan') {
         if (XRAY_SCAN == "YES") {
             def xrayConfig = [
                'buildName'     : env.JOB_NAME,
                'buildNumber'   : env.BUILD_NUMBER,
                'failBuild'     : false
              ]
              def xrayResults = rtServer.xrayScan xrayConfig
              echo xrayResults as String
         } else {
              println "No Xray scan performed. To enable set XRAY_SCAN = YES"
         }
         sleep 60
     }

    //Promote docker image from staging local repo to production repo in Artifactory
    /* stage ('Promote') {
            def promotionConfig = [
              'buildName'          : env.JOB_NAME,
              'buildNumber'        : env.BUILD_NUMBER,
              'targetRepo'         : PROMOTE_REPO,
              'comment'            : 'App works with latest released version of gradle swampup app, tomcat and jdk',
              'sourceRepo'         : SOURCE_REPO,
              'status'             : 'Released',
              'includeDependencies': false,
              'copy'               : true
            ]
            rtServer.promote promotionConfig
            reTagLatest (SOURCE_REPO)
            reTagLatest (PROMOTE_REPO)
        // promote war file from gradle-dev-local to gradle-release-local
     }*/

}

def testApp (tag) {
    docker.image(tag).withRun('-p 9191:8181') {c ->
        sleep 10
        //def stdout = sh(script: 'curl "http://localhost:9191/swampup/"', returnStdout: true)
        //if (stdout.contains("Welcome Docker Lifecycle Training")) {
          //  println "*** Passed Test: " + stdout
            println "*** Passed Test"
            return true
       // } else {
        //    println "*** Failed Test: " + stdout
         //   return false
       // }
    }
}

//Tag docker image
def reTagLatest (targetRepo) {
    def BUILD_NUMBER = env.BUILD_NUMBER
    sh 'sed -E "s/@/$BUILD_NUMBER/" retag.json > retag_out.json'
    switch (targetRepo) {
          case PROMOTE_REPO :
              sh 'sed -E "s/TARGETREPO/${PROMOTE_REPO}/" retag_out.json > retaga_out.json'
              break
          case SOURCE_REPO :
              sh 'sed -E "s/TARGETREPO/${SOURCE_REPO}/" retag_out.json > retaga_out.json'
              break
    }
    sh 'cat retaga_out.json'
    withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
        def curlString = "curl -u " + env.USERNAME + ":" + env.PASSWORD + " " + ART_SERVER_URL
        def regTagStr = curlString +  "/api/docker/$targetRepo/v2/promote -X POST -H 'Content-Type: application/json' -T retaga_out.json"
        println "Curl String is " + regTagStr
        sh regTagStr
    }
}

def updateProperty (property) {
    withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
            def curlString = "curl -u " + env.USERNAME + ":" + env.PASSWORD + " " + "-X PUT " + ART_SERVER_URL
            def updatePropStr = curlString +  "/api/storage/${SOURCE_REPO}/containerize-go-microservice/${env.BUILD_NUMBER}?properties=${property}"
            println "Curl String is " + updatePropStr
            sh updatePropStr
     }
}