import subprocess

def lambda_handler(event, context):
    # place executable created from go build command in the same directory as this file
    # zip that file and this one together (not in a folder)
    # then upload to aws lambda function
    args = "./spotify-discover-yearly"
    
    popen = subprocess.Popen(args, stdout=subprocess.PIPE)
    popen.wait()
    
    output = popen.stdout.read()
    print(output)