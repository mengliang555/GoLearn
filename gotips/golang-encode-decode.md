# Encode and Decode
- Encode -> make the struct(object) to array byte. So we can endurance or transport the data.
- Decode ->  Conversely Encode. 

## Json & GOB & Protobuf
- [StackOverFlow: Difference between encoding/gob and encoding/json](https://stackoverflow.com/questions/41179453/difference-between-encoding-gob-and-encoding-json)
- When using json.Marshal(struct)
   - If the struct's Attributes(I'm not sure the definition) has no json-tag. The json's key will be the Attributes original name
   ```
   //with no json tag
   type HelloWorld struct{
       Hello string 
       World string
   }
   // will be {"Hello":"","World":""}

   // with the json tag. the json-key will be the json tag
   type HelloWorld struct{
       Hello string `json:"hello"` 
       World string `json:"hello"`
   }
   // will be {"hello":"","world":""}
   ``` 
   - **If you haven't fully test you code, never add or change the tag. Consider your old data please**