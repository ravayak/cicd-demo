node {
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
        }
        stage("build"){
            sh "go build -v -o app ."
            sh "git add ."
            sh "git commit -m 'built app'"
        }
    }
}