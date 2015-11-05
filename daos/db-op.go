package daos

import (
    //"fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "log"
    "sync"
    "time"
    
)

const (
    MongoDBHosts = "localhost:27017"
    AuthDatabase = "posts"
    AuthUserName = "sc_admin"
    AuthPassword = "sc_admin"
    TestDatabase = "posts"
)


type (
	// PostContent contains information for an individual content.
	PostContent struct {
		Title         string    `bson:"title"`
		Link          string    `bson:"link"`
        Name          string    `bson:"name"`
		Comment       string    `bson:"comment"`
	}


	// Posts contains information for an individual post message.
	Posts struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		UserId    string        `bson:"user_id"`
        Active    bool          `bson:"active"`
		Content   PostContent   `bson:"content"`
        CreatedAt string        `bson:"created_at"`
	}
)

type Person struct {
        Name string
        Phone string
}



// main is the entry point for the application.
func main() {
    // We need this object to establish a session to our MongoDB
    mongoDBDialInfo := &mgo.DialInfo {
        Addrs:  []string{MongoDBHosts},
        Timeout: 60* time.Second,
        Database: AuthDatabase,
        Username: AuthUserName,
        Password: AuthPassword,
    }
    
    // Create a session which maintains a pool of scoket connections
    // to our MongoDB
    mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
    if err != nil {
        
        log.Fatalf("CreateSession: %s\n", err)
    } 
    
     // Optional. Switch the session to a monotonic behavior.
    mongoSession.SetMode(mgo.Monotonic, true)
    
    /* verify the db connection and insert operation */
    /*
    c := mongoSession.DB("test").C("people")
         err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
 	               &Person{"Cla", "+55 53 8402 8510"})
         if err != nil {
                 log.Fatal(err)
         }

         result := Person{}
         err = c.Find(bson.M{"name": "Ale"}).One(&result)
         if err != nil {
                 log.Fatal(err)
         }

         fmt.Println("Phone:", result.Phone)
    */
    
    // Create a wait group to mange the goroutines.
    var waitGroup sync.WaitGroup

    // Perform 10 concurrent queries against the database.
    waitGroup.Add(10)
    for query :=0; query < 10; query++ {
        go RunQuery(query, &waitGroup, mongoSession)
    }

    // Wait for all the queries to complete.
    waitGroup.Wait()
    log.Println("All Queries Completed")
}




// RunQuery is a function that is lauched as a goroutine to perform 
// the MongoDB work.

func RunQuery(query int, waitGroup *sync.WaitGroup, mongoSession *mgo.Session) {
    // Decrement the wait group count so the program knows this 
    // has been completed once the goroutine exists.
    defer waitGroup.Done()
    
    //Request a socket connection from the session to progress our query.
    // Close the session when the goroutine exits and put the connection back
    // into the pool.
    sessionCopy := mongoSession.Copy()
    defer sessionCopy.Close()
    
    // Get a collection to execute the query against.
    collection := sessionCopy.DB(TestDatabase).C("posts")
    
    log.Printf("RunQuery: %d : Executing\n", query)
    
    // Retreive the list of posts.
    var posts [] Posts
    err := collection.Find(nil).All(&posts)
    
    if err != nil {
        log.Printf("RunQuery: ERROR : %s\n", err)
        return 
    }
    log.Println(posts)
    log.Printf("RunQuery : %d : Count[%d]\n", query, len(posts))
}
