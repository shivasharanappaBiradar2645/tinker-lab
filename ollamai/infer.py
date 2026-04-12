import ollama

client = ollama.Client(host="http://10.117.63.44:11434")

response = client.chat(  
model="gemma3:270M",  
messages=[  
{"role": "user", "content": "Explain quantum computing simply"}  
],
options = {
"num_predict":50}  
)

print(response["message"]["content"])


response = client.generate(
    model="gemma3:270m",
    prompt="Write a haiku about programming"
)

print(response["response"])

ctx = response["context"]
response = client.generate(
	model="gemma3:270m",
	prompt="its good keep it up",
	context = ctx)

print(response["response"])


stream = client.generate(
    model="gemma3:270m",
    prompt="Explain transformers in ML",
    stream=True
)

for chunk in stream:
    print(chunk["response"], end="", flush=True)


client.list()
