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

    `sorry`: {
        "No need to say sorry. This is a safe place.",
    },
    `i need (.*)`: {
        "Why do you need %s?",
    },
    `i can't (.*)`: {
    "You can do anything is you try hard enough.",
    "Why can't you %s?",
    },
    `because (.*)`: {
        "because why?",
        "you seem to be avoiding my questions",	
    },
    `i am (.*)`: {
        "You are %s? Why are you %s?",
    },
    `I think (.*)`: {
        "Why do you think that?",
    },
    `i feel (.*)`: {
        "Why do you feel like that?",
    },
    `i have (.*)`: {
    "Why have you %s?",
    },
    `i would (.*)`: {
        "Why would you %s?",
    },
    `i don't (.*)`: {
        "Why don't you %s?",
    },

    // Questions
    `why don't you (.*)`: {
        "That's a hard one to answer.",        
    },
    `are you (.*)`: {
    "I'm only trying to help.",
    },
    `is it (.*)`: {
        "How do you know it is %s?",
        "is it %s?",
    },
    `can you (.*)`: {
        "What makes you think I can %s?",
    },

    `is there (.*)`: {
        "Is there %s? I don't understand?",
    },
    `why (.*)`: {
        "Why are you asking that?",
        "Why is %s important?",
    },
    `what (.*)`: {
        "Why are you asking me 'what %s'",
    },
    `can I (.*)`: {
        "Can you %s?",
    },
    `(.*)\?`: {
        "Thats an strange question?",
        "Could you rephase that?",
    },
    }

    // Failing to find an answer in the libary above this libary will be used
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
var Similar = map[string][]string{
    // Expressions
    "want":     []string{"need", "require", "demand"},
    "i am":     []string{"i'm"},
    "you are":  []string{"you're"},
    "sorry":    []string{"regretful","apologies","excuse me"},
    "hello":    []string{"hey","hi","greetings","howya"},
    "thank you":[]string{"thanks","bless","praise"},
    "happy":    []string{"good"},  
    "father":   []string{"dad","daddy"},
    "mother":   []string{"mam","mammy"},
    "sad":      []string{"bad", "upset"},
}