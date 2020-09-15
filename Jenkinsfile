pipeline {
   agent any
   
   tools {
      go 'Go1.15'
   }
   
   environment {
      GO115MODULE = 'on'
      CGO_ENABLED = 0 
      GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
   }
   
   stages {
      stage('Pull') {
         steps{
            git 'https://github.com/nickormushev/Jenkins_test.git'
         }
      }
   
      stage("Pre_test") {
         steps {
            echo 'Installing dependencies'
            sh 'go version'
            sh 'go get -u golang.org/x/lint/golint'
         }
      }
   
      stage("Build") {
         steps {
            echo 'Building project'
            sh 'go build'
         }
      }
   
      stage("Test") {
         steps {
            withEnv(["PATH+GO=${GOPATH}/bin"]){
               echo 'Running vetting'
               sh 'go vet .'
               echo 'Running linting'
               sh 'golint .'
               echo 'Running test'
               sh 'go test -v'
            }
         }
      }
   }
   
   post {
      always {
         echo "Job has finished"
      }
   }
}
