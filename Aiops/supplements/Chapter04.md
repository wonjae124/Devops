# 0. Aiops Roadmap


# 1. Github Action-CI/CD
- go언어 파일의 테스트 진행
- import `"testing" `
    - `main.go`
    ```golang
          package main

          import "fmt"

          func main() {
              msg := sayHello("Alice")
              fmt.Println(msg)
          }

          func sayHello(name string) string {
              return fmt.Sprintf("Hi %s", name)
          }
    
    ```
    - `main_test.go`
        ```golang
            package main

            import "testing"

            func Test_sayHello(t *testing.T) {
                name := "Bob"
                want := "Hello Bob"

                if got := sayHello(name); got != want {
                    t.Errorf("hello() = %q, want %q", got, want)
                }
            }

        ```
     - `go test .`   
        
- <img src = "https://github.com/wonjae124/Devops/blob/main/image/%EC%8A%A4%ED%81%AC%EB%A6%B0%EC%83%B7%202023-03-17%2018-27-38.png" width=1200>

# 느낀점
- 다음부터, 웬만한 작업에 테스트 코드를 만들어야겠다.
- 

#### 출처
- [Continuous integration with Go and GitHub Actions]([https://youtu.be/MpYB4Qcl570 ](https://www.alexedwards.net/blog/ci-with-go-and-github-actions))
<br/><br/><br/>
