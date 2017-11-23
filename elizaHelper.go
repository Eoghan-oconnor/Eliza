package eliza

// Imports 
import (
    "fmt"
    "math/rand"
    "regexp"
    "strings"
)

// replys to user by restructing the user input 
func Reply(input string) string {

    //Passes user input into the preprocess function    
    input = preprocess(input)

    for pattern, responses := range phrases {
        re := regexp.MustCompile(pattern)
        matches := re.FindStringSubmatch(input)

        if len(matches) > 0 {
            var fragment string
            if len(matches) > 1 {
                fragment = reflect(matches[1])
            }

            //this selects a random reponse based on the keywords 
            output := randChoice(responses)
            
            //Puts input and output together so it appears smart 
            if strings.Contains(output, "%s") {
                output = fmt.Sprintf(output, fragment)
            }
            return output
        }
    }

    //if no response is found it selects randomly from convoRestart
    return randChoice(convoRestart)
}

//Allows user input to be 'understood by eliza'
func preprocess(input string) string {
    input = strings.TrimRight(input, "\n.!")
    input = strings.ToLower(input)

    formattedInput := strings.Split(input, " ")
	for i, word := range formattedInput {
		formattedInput[i] = strings.ToLower(strings.Trim(word, ".! \n"))
    }
    
    formattedInput = PostProcess(formattedInput)

    input = strings.Join(formattedInput," ")

    return input
}

//Searches Reflections to make responses about the user
func reflect(fragment string) string {
    words := strings.Split(fragment, " ")
    for i, word := range words {
        if Reflections, ok := Reflections[word]; ok {
            words[i] = Reflections
        }
    }
    return strings.Join(words, " ")
}

// PostProcess function loops threw the Similar to 'dumb down' words that are similar
// This wwill make the bot look and feel more intellegent, giving it an illusion of understanding
func PostProcess(parsed []string) (input []string) {

	search := make(map[string]string)
	for key, list := range Similar {
		for _, word := range list {
			search[word] = key
		}
	}

	input = parsed
	for i, word := range parsed {
		if processedWord, ok := search[word]; ok {
			input[i] = processedWord
		}
	}
	return input
}

// Randomizes runmber for lists in elizas scripts
func randChoice(list []string) string {
    randIndex := rand.Intn(len(list))
    return list[randIndex]
}