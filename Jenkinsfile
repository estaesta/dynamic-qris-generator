pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'docker build -t $DOCKER_IMAGE .'
      }
    }

    stage('Deploy Docker') {
      steps {
        sh '''docker stop qris-api || true
docker rm qris-api || true
docker run -d --name qris-api -p $PORT:8080 -e PORT=8080 $DOCKER_IMAGE'''
      }
    }

  }
  environment {
    DOCKER_IMAGE = 'qris-api:latest'
    PORT = '8888'
  }
}