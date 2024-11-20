import Foundation


let jsonString1 = """
{
    "name": "John Doe",
    "age": 30,
    "email": "johndoe@example.com"
}
"""
let jsonString2 = """
{
    "name": "Jane Doe",
    "email": "janedoe@example.com"
}
"""
let jsonString3 = "Invalid JSON"

let user1 = unmarshalUser(jsonString: jsonString1)
let user2 = unmarshalUser(jsonString: jsonString2)
let user3 = unmarshalUser(jsonString: jsonString3)

assert(user1?.name == "John Doe")
assert(user1?.age == 30)
assert(user1?.email == "johndoe@example.com")

assert(user2?.name == "Jane Doe")
assert(user2?.age == nil)
assert(user2?.email == "janedoe@example.com")

assert(user3 == nil)