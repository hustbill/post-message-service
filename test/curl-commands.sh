curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Bob Smith", "gender": "male", "age": 50}' http://127.0.0.1:3000/user
{"id":"5497246c380a967ff1000003","name":"Bob Smith","gender":"male","age":50}



curl -H "Content-Type: application/json" -X POST -d '{"userId": 1192293, "orderNumber": "J2326", "orderTotal": 34.45, "orderDate": "2015-09-30T07:00:00.000Z", "orderState": "paid" }' -v http://127.0.0.1:8098/v1/orders
echo "Test case 4 Done "



 #Insert a new post  
 curl -XPOST -H 'Content-Type: application/json' -d '{"user-id": 101, "type": "text", "content": "Hello World!"}' http://127.0.0.1:3000/v1/posts 
Result: {"id":"563aa288d1261946cb000001","type":"text","content":"Hello World!","user-id":101}

 #Query an existing post
curl -H "Content-Type: application/json" -X GET -v http://127.0.0.1:3000/v1/posts/563aa288d1261946cb000001


 # Insert a new product 
 curl -XPOST -H 'Content-Type: application/json' -d '{"Name": "peanuts", "Description": "Honey Roasted peanuts", "MetaDescription": "Good taste food", "Permalink": "www.peanuts.com"}' http://127.0.0.1:3000/v1/products 

 curl -XPOST -H 'Content-Type: application/json' -d '{"Name": "boil egg", "Description": "Safeway signure egg", "MetaDescription": " best by 10 days", "Permalink": "www.goodegg.net"}' http://127.0.0.1:3000/v1/products 



 # Query an existing product
curl -H "Content-Type: application/json" -X GET -v http://127.0.0.1:3000/v1/products/1

 # Query an existing product with q (search string)
curl -H "Content-Type: application/json" -X GET -v http://127.0.0.1:3000/v1/products/1 -d '{"q" : "boil egg"}'