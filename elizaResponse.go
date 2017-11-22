package eliza

// Libary of phrases
var phrases = map[string][]string{
    // keywords and responses
    `hello`: {
        "Hello How are you feeling?",   

    },
    `sad`: {
        "Why are you sad?",
        "What's making you sad?",
        "Tell me more about why your sad?",
    },
   
    `today`: {
        "Tell me more about today",
        "What happened today?",
    },
    `father`: {
        "Tell about your father?",
    },
    `mother`: {
        "tell me about your Mother?",
        
    },
    `family`: {
        "If I mention your family, how does that make you feel?",
    },
    `happy`: {
        "Why are you happy?",
        "I'm glad you're happy but what made you happy?",
    },
}  

var convoRestart= []string{
    "I'm sorry but i dont understand what you're saying",
    "Let's focus on whats important.... how are you feeling",
	"Tell me more.",
}

//changes the pronoun to allow conversations to make sense
var Reflections = map[string]string{
    "am":       "are",
    "was":      "were",
    "I":        "you",
    "I'd":      "you would",
    "I've":     "you have",
    "I'll":     "you will",
    "my":       "your",
    "are":      "am",
    "you've":   "I have",
    "you'll":   "I will",
    "you're":   "I'm",
    "your":     "my",
    "yours":    "mine",
    "you":      "I",
    "me":       "you",
}

//Allows user to enter words similar and have it still make sense
var Synonymizer = map[string][]string{
    // Expressions
   
    "hello":    []string{"hey","hi","greetings","howya"},
    
}