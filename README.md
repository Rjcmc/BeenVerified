Installation based on you are using linux:
	First you need to install golang on your computer, you can follow the instruction on the official website https://golang.org 
	The second step is install goji, you need be on the project folder and open a terminal, on the terminal you shuld type the following command: "go get goji.io"
	Finally you will need to install de sqlite database,open a terminal and put the next command "sudo apt-get install sqlite3"
	
For using this API you need open a terminal on the project folder and type this:"go run main.go"
open a browser and in the search bar paste one of the followings URLs
	http://localhost:8000/artist/Santana  		//this one is for search by artist,instead of santana you should put the name of the artis 
	http://localhost:8000/song/Gala			//this one is for search by song,instead of Gala you should put the name of the song
	http://localhost:8000/genre/Classic Rock	//this one is for search by genre,instead of Classic Rock you should put the name of the genre
	http://localhost:8000/order  			//this one is to get genres with the count of their songs and the sum of their lenght
	http://localhost:8000/length/1/167		//this one is for search by lenght,instead of 1/167 you should put the min and the max that 								you are looking for
	
