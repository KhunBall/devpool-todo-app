pipeline {
    agent any

    environment {
		DOCKERHUB_CREDENTIALS=credentials('1447b088-c586-4638-9056-1ac7a24db4cf')
	}

    stages {
        stage('Clean old container if exist') {
            steps {
                sh 'docker ps -q -f status=exited | xargs --no-run-if-empty docker rm'
                sh 'docker rm -f docker-todo-app'
            }
        }
        stage('Clean old image if exist') {
            steps {
                sh 'docker image prune --filter="dangling=true"'
                // sh 'docker images -q -f status=exited | xargs --no-run-if-empty docker rmi'
                // sh 'docker rmi khunball/docker-todo-app:latest'
            }
        }
        stage('Deploy to server') {
            steps {
                    sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                    sh 'docker build --tag khunball/docker-todo-app .'
                    sh 'docker push khunball/docker-todo-app'
            }
        }
        stage('Test run') {
            steps {
                sh 'docker run --name docker-todo-app -dp 80:80 khunball/docker-todo-app'
            }
        }       
        stage('Deploy') {
            steps {
                    sshagent(['prod-credential']) {
                    sh 'scp -o StrictHostKeyChecking=no docker-compose.yml khunball@192.168.1.23:/home/khunball/docker-compose.yml'
                    sh 'ssh -o StrictHostKeyChecking=no khunball@192.168.1.23 docker compose up -d'
                }
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