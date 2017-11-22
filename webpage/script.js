//input variable
const input = $("#userInput");
//output variables
const ouputText = $("#output");

// user input keypress
input.keypress(function(e){
    if(e.keyCode != 13){
        return;
    }

    //this stops the page from refreshing allowing the coversation to flow
    e.preventDefault();

    // Variable 
    const userInput = input.val();
    input.val("");

    const user = {"input" : userInput }
    ouputText.append( "<p class=\"reply\">" + "<b class=\"userName\">User: </b>" + userInput + "</p>");
     
    $.get("/elizaPrompt", user)
        .done(response => {
            //creating the output
            const output = "<p class=\"reply\">" + "<b class=\"userName\">Eliza: </b>" + response + "</p>";

            //Eliza responses are outputted to the webpage 
            setTimeout(function(){ouputText.append(output)});
            
        })
});