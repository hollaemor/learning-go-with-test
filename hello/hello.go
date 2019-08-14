package main

import "fmt"

const (

		englishPrefix = "Hello, "
		spanish = "Spanish"
		spanishHelloPrefix = "Hola, "
		french = "French"
		frenchHelloPrefix = "Bonjour, "
)


func greetingPrefix(language string) (prefix string) {
		   switch language {
		  case french:
					 prefix = frenchHelloPrefix
		case spanish:
					 prefix = spanishHelloPrefix
		default:
				prefix = englishPrefix
		}
return

}
func Hello(name, language string) string {
  if name == ""  {
		name = "World"
  }

  return greetingPrefix(language)  + name
}


func main() {
 fmt.Println(Hello("world", "") )
}
