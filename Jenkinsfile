pipeline {
  agent any
  stages {
    stage('Stop & Clean') {
      steps {
        dir(path: '/srv/blockchain/first-network')
        sh './byfn.sh down'
        sh 'docker rmi -f $(docker images | grep fabcar | awk \'{print $3}\')'
      }
    }
    stage('Pull') {
      steps {
        dir(path: '/srv/blockchain')
        sh 'git pull . master:master'
      }
    }
    stage('Start Fabric') {
      steps {
        dir(path: '/srv/blockchain/fabcar')
        sh './startFabric.sh'
      }
    }
  }
}