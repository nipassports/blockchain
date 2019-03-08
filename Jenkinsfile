pipeline {
  agent any
  stages {
    stage('Stop Fabric') {
      parallel {
        stage('Stop') {
          steps {
            sh 'docker rm -f $(docker ps | grep "hyperledger/fabric" | cut -d" " -f1) | exit 0'
            sh 'docker rm -f $(docker ps | grep "dev-peer" | cut -d" " -f1) | exit 0'
          }
        }
        stage('Clean') {
          steps {
            sh 'docker rmi -f $(docker images | grep "hyperledger/fabric" | awk \'{print $3}\') | exit 0'
            sh 'docker rmi -f $(docker images | grep fabcar | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep "hyperledger/fabric" | cut -d" " -f1) | exit 0'
            sh 'docker rmi -f $(docker images | grep "dev-peer" | cut -d" " -f1) | exit 0'
          }
        }
      }
    }
    stage('Build') {
      steps {
        sh './create_bin.sh'
      }
    }
    stage('Start Fabric') {
      steps {
        sh 'cd fabcar && ./startFabric.sh'
        sh 'cd fabcar && ./actionsPostStart.sh'
      }
    }
  }
}
