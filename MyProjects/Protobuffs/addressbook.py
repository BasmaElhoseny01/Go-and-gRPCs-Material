import addressbook_pb2
addressbook = addressbook_pb2.AddressBook()

for i in range(10):
    person = addressbook.people.add()
    person.id=i
    person.name="7amada"
    person.email="7amada@gmail.com"

    # add phone to the list of the person's phone list
    phone=person.phones.add()
    phone.number="123123"
    phone.type = addressbook_pb2.Person.PHONE_TYPE_HOME


print(addressbook)
print(addressbook.SerializeToString())

with open('proto.out','wb')as f:
    f.write(addressbook.SerializeToString())

import json
persons=[]
for i in range(10):
    persons.append({
        "id":i,
        "name":"7amada",
        "email":"7amada@gmail.com",
        "phones":[{"number":"123123",'type':"HOME"}]
    })

print(persons)
with open('json.out','w')as f:
    f.write(json.dumps(persons))


# PS D:\Go Sockets & RPCs\Protobuffs> dir *.out
    
# Mode                 LastWriteTime         Length Name
# ----                 -------------         ------ ----
# -a----         3/13/2024   3:25 PM           1080 json.out
# -a----         3/13/2024   3:25 PM            418 proto.out

# 40% less size :D
