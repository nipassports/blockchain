pipeline {
  agent any
  stages {
    stage('Stop Fabric') {
      steps {
        sh 'docker rm -f $(docker ps | grep nip -v | cut -d" " -f1 | tail -n +2) | exit 0'
        sh 'docker rmi -f $(docker images | grep fabcar | awk \'{print $3}\') | exit 0'
      }
    }
    stage('Start Fabric') {
      steps {
        dir(path: 'fabcar')
        sh './startFabric.sh'
      }
    }
  }
}