node {
    def root = tool type: 'go', name: '1.20'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stage("clone"){
             checkout scm 
        }
        sh 'go version'
        stage("build"){
            go build -v -o app .
        }
    }
}