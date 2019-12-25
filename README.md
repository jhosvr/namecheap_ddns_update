# namecheap_ddns_update

This is mostly intended for homelabbers who use Namecheap as a domain registration service. 

It checks for the client's IP address by using an IP service (list of urls), and then updates the Namecheap A record for that address using Namecheap's Dynamic DNS service.

# Using namecheap_ddns_update

1. Fill in the user constants in `main.go`. I opted to just bake these values into source in order to stay as simple as possible.
2. Build the package 


# References
* [Namecheap - how to use Dynamic DNS](https://www.namecheap.com/support/knowledgebase/article.aspx/36/11/how-do-i-start-using-dynamic-dns)
* [Namecheap - how to use a web request to update dynamic DNS](https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-do-i-use-a-browser-to-dynamically-update-the-hosts-ip)
