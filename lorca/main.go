package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
    	<meta charset="UTF-8">
    	<meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login | Controle de Coletor</title>
   		<!-- <link rel="stylesheet" type="text/css" href="/assets/css/main.css"> -->
    <link rel="stylesheet" type="text/css" href="/assets/css/partials/page_1.css">
	</head>
		<body>
    	<main class="container">
        	<h2>Login</h2>
        <form action="" id="login">
            	<div class="input-field">
                	<input type="email" name="username" id="email"
                    	placeholder="Entre com seu e-mail" required="required">
                	<div class="underline"></div>
            	</div>
            	<div>
            	</div>
            	<div class="input-field">
                	<input type="password" name="password" id="senha"
                    	placeholder="Entre com a senha" required="required">
                	<div class="underline"></div>
            	</div>
            	<input type="submit" value="Entrar" id ="logando"> 
        	</form>
        	<br> </br>
        		<div class="footer">
            	<span>© J.VI.S</span>
                	</div>
            	</div>

        	</div>
    	</main>
    		<script src="/assets/js/jquery.min.js"></script>
    		<script src="/assets/js/login.js"></script>
		</body>
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
