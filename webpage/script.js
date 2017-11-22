//input variable
const inputText = $("#userInput");
//output variables
const ouputText = $("#botOutput");

// user input keypress
inputText.keypress(function(e){

    if(e.keyCode != 13){
        return;
    }

    //this stops the page from refreshing allowing the coversation to flow
    e.preventDefault();

    // Variable 
    const userInput = inputText.val();
    inputText.val("");
    const userPrompt = {"input" : userInput }
    ouputText.append( "<p class=\"reply\">" + "<b class=\"userName\">User: </b>" + userInput + "</p>");
    
    // When go can comunicate with webpage '.done' will run otherwise '.fail' will
    $.get("/elizaPrompt", userPrompt)
        .done(response => {
            // if the webpage can connect to the GO eliza files then the user will be prompted with a response from eliza
            const output = "<p class=\"reply\">" + "<b class=\"userName\">Eliza: </b>" + response + "</p>";

            // The bots response with HTML elements are output to the webpage with an alert
            setTimeout(function(){ouputText.append(output), new Audio('assets/sound/responseAlert.mp3').play();}, 1000);
            
        })
});