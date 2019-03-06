pipeline {
  agent any
  stages {
    stage('Stop Fabric') {
      parallel {
        stage('Stop') {
          steps {
            sh 'docker rm -f $(docker ps | grep nip -v | cut -d" " -f1 | tail -n +2) | exit 0'
            sh 'docker rmi -f $(docker images | grep fabcar | awk \'{print $3}\') | exit 0'
          }
        }
        stage('Clean') {
          steps {
            sh 'docker rmi -f $(docker images | grep hyperledger | awk \'{print $3}\') | exit 0'
            sh 'docker rmi -f $(docker images | grep hyperledger | awk \'{print $3}\') | exit 0'
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
