# Azure Storage Put Block from URL Escaping Issue Demo/Debug

Question on Azure Q&A: https://docs.microsoft.com/en-us/answers/questions/641723/i-can39t-get-azure-storage-to-support-putting-data.html

Client Counterpart: https://github.com/nelsonjchen/PutBlockFromUrlEscIssueDemoClient

# Usage

Build the Dockerfile or Go program and host it somewhere public.

Edit the URL in the Client Counterpart to point to the public host.

Run the client and it will throw an exception as Azure Storage escapes a necessary "`%2F`" in the URL. 

# Demo Host

Here is an example of a URL that cannot be PUT to Azure Storage:

https://put-block-from-url-esc-issue-demo-server-3vngqvvpoq-uc.a.run.app/red%2Fblue.txt

