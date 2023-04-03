executable : main.o read_and_write.o 
		gcc -o executable main.o read_and_write.o -lcurl
	
main.o : main.c
		gcc -c $<

read_and_write.o : read_and_write.c biblio.h
		gcc -c $<

clean :
		rm  -rf *.o