pipeline {
    agent any
    stages {
        stage('Clean old container if exist') {
            steps {
                sh 'docker ps -q -f status=exited | xargs --no-run-if-empty docker rm'
                sh 'docker rm -f docker-todo-app'
            }
        }
        stage('Clean old image if exist') {
            steps {
                sh 'docker images -q -f dangling=true | xargs --no-run-if-empty docker rmi'
                sh 'docker image rm khunball/docker-todo-app:latest'
            }
        }
        stage('Deploy to server') {
            steps {
                sh 'docker build --tag khunball/docker-todo-app .'
                sh 'docker push khunball/docker-todo-app'
            }
        }
        stage('Test run') {
            steps {
                sh 'docker run --name docker-todo-app -dp 80:80 khunball/docker-todo-app'
                sh 'docker stop docker-todo-app'
            }
        }       
        stage('Deploy') {
            steps {
                sh 'scp -o StrictHostKeyChecking=no docker-compose.yml khunball@192.168.1.23:/home/khunball/docker-compose.yml'
                sh 'ssh -o StrictHostKeyChecking=no khunball@192.168.1.23 docker-compose up -d'
            }
        }
        stage('Clean up') {
            steps {
                sh 'docker rm -f docker-todo-app'
                sh 'docker image rm khunball/docker-todo-app:latest'
            }
        }
    }
}