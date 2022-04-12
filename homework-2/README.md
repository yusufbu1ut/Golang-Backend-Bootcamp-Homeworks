# homework-2 / yusufbu1ut

## main.go

In here firstly, created init func to fill the slice of books. Authors and books created after builded with NewBook and NewAuthor constructor funcs. Lastly added into slice.

After slice adding processes, program takes arguments there is for subcommand 'list', 'search', 'buy' and 'delete'(List command added to check slice items changings ). 
List command prints all not deleted slice items. Search command searchs in slice what it comes after search command(it can be a book name, writer name or surname, 
stock code or stock number). Buy command takes two argument one id, one count. These arguments convert to int and goes to process in book.go Buy func with count parameter. 
Delete command takes one int argument. In this case, program uses func named Delete which takes deletable parameter(Book) and processes delete command.

ps: All error prints given in err.go . (Extra value, invalid value etc.) 


## helper

* funcs.go
  
  In here there is two func Search and List; List func prints all not deleted book items. Search prints in Book slice for given input, searching can be with book name, writer name or surname
  ,exact stock code or stock number. Compare processes runs with toLower compares for book name and writer, sku processes runs with str to int converting for given count.


## models

* book.go
  
  In here there is two struct Author and Book, Author embedded in Book struct. These structs has NewBook and NewAuthor constructor funcs to identify. In Book struct: pages(int), price(float),
  stock number(int), ISBN(int) created with RandomInt and RandFloat funcs;Id and stock code created with global counts in book.go when creating new book item these counts increasing
  one point. Book Init constructor func takes two parameter book name and Author. Author has name,surname and isDeleted(to try something) fields. Author NewAuthor constructor takes two value name 
  and surname. Author Print func is to use on printing and comparing processes.
  
  Buy method takes one int parameter checks stock number with this parameter and if it is not higher  than stock number decreases stock number.
  
  There is deletable inteface contains delete func. This func is calling when item.delete func called and changes book items IsDeleted to true.

* generators.go

  In here there is func RandomInt and RandFloat takes two int parameter returns int value and float value.
