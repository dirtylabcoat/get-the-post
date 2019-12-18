# go-get-the-post
Proxies POST, PUT and DELETE calls through GET.

Deploy with:

	gcloud functions deploy getthepost --runtime go111 --allow-unauthenticated --trigger-http --entry-point GetThePost --region europe-west2

Test with:

	curl -X POST -d'{"url":"https://www.google.se","method":"POST","data":"bla"}' https://REGION-PROJECT.cloudfunctions.net/getthepost


