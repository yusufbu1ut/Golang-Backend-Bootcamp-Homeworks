# Homework-3 / yusufbu1ut
## About Work 
***

Firstly, creating entity elements in internal/sample/csv_helper.go, after this process program connections executes with DB and executing migrations, lastly adding these items on DB with models InsertSampleData(model) funcs.

There is 4 command list,search,buy and delete;

- list   
List an use like just 'list' command returns all books.     
Can use like 'list b' with 'b' sub command returns all books.   
Can use like 'list a' with 'a' sub command retuns all authors.

- search    
Search command takes input parameter for example 'search harry potter'. In this command input parameter firstly checks in books names, if there is no book such as with that name, after it checks in authors name_surname. If input is in books returns books and books' authors, else retuns authors and authors' books.

- buy    
Buy command takes 2 positive int count as parameters, first is book' id that wanted to buy, other one is count that how many do you want to buy. Checks id if there is a book with same id, after checks given count and process on stock_amount.

- delete    
Delete command takes just one parameter that is book id represents which one do you want to delete. If there is , program finds it deletes from db. In this process, deleting processes working on books, book_authors which as same isbn number with deleted book and authors that auhtor has no not deleted books.

#### PS: Added Create() and Update() repository funcs in repositories, but not used.


## Database
---
In this work, I used PostgreSQL.

### books
***
ID         
CreatedAt   
UpdatedAt   
DeletedAt   
Name
Pages       
Publisher   
StockCode   //random int    
StockAmount     //random int   
ISBN    (unique count and foreing key with referance book_authors BookID)    
Price   //random float       

### book_authors
***
ID         
CreatedAt   
UpdatedAt   
DeletedAt   
BookID   (foreing key with referance books ISBN)  
AuthorID    (foreing key with referance authors ID)  

### authors
***
ID         
CreatedAt   
UpdatedAt   
DeletedAt   
NameSurname     
Age     //random int        

## Information
internal/helper has two different way to add datas from csv. First, /csvToDB creates datas and adds directly in to the db. Second, /readInsert creates datas bring them together in a slice after adds to db with repository.InsertSampleData funcs.  