node {
    //  Jenkins utilise automatiquement l'agent par défaut "any" si on spécifie rien

    //  la fonction tool permet de définir des configurations d'outils dans Jenkins
    //  elle permet de déclarer et d'utiliser des outils externes tels que les SDK, les compilateurs, les interpréteurs, etc.
    // using the name defined in the Global Tool Configuration (1.20)
    def root = tool type: 'go', name: '1.20'

    // Voir doc plugin golang jenkins: https://plugins.jenkins.io/golang/
    // Export environment variables to pointing the Go installation;
    // the `PATH+X` syntax prepends an item to the existing `PATH`:
    // https://jenkins.io/doc/pipeline/steps/workflow-basic-steps/#withenv-set-environment-variables  
    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
        }
        stage("build"){
            // git part

            // global config
            sh 'git config --global user.email "antonny.kihm@protonmail.com"'
            sh 'git config --global user.name "ravayak"'
            sh "go build -v -o app ."

            // Ne pas oublier de rajouter le binaire créé            
            sh "git add ."
            sh "git commit -m 'built app'"
            // Création du tag
            sh "git tag -l build-jenkins"   
            sh "git tag -a -f -m 'Jenkins build' build-jenkins"
            sh "git --version"

            // docker section

            // build de l'image
            sh 'docker build . -t ynno/go-cicd-demo -f DockerFile-build'
        }
        stage('deliver') {
                withCredentials([usernamePassword(
                    credentialsId: 'GITHUB_TOKEN',
                    passwordVariable: 'TOKEN',
                    usernameVariable: 'USER')]) {
                    sh 'git push https://${USER}:${TOKEN}@github.com/ravayak/cicd-demo.git build-jenkins'
            }

            withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
                sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
                sh 'docker push ynno/go-cicd-demo'
            }

            // exemple de push sur un registry local créé avec nexus : https://www.sonatype.com/products/sonatype-nexus-repository
            withCredentials([usernamePassword(credentialsId: 'dockerhublocal', passwordVariable: 'dockerhublocalPassword', usernameVariable: 'dockerhublocalUser')]) {
                sh "docker tag ynno/go-cicd-demo 9532-2-4-250-10.ngrok-free.app/go-cicd-demo:latest"
                sh "docker login 9532-2-4-250-10.ngrok-free.app -u ${env.dockerhublocalUser} -p ${env.dockerhublocalPassword}"
                sh 'docker push 9532-2-4-250-10.ngrok-free.app/go-cicd-demo:latest'
            }

        }
    }
}