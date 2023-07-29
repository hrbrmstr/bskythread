# Bluesky Thread Retriever Gateway/Proxy

A simple Go application to interact with bsky.social API. 

Retrieving post/thread data requires a Bluesky login, which is still a pretty rare thing. This code acts as a gateway/proxy, enabling post/thread retrieval with just a `did:plc` and post id/rkey.

You'll need a valid bsky user id and app password to be able to start the gateway.

One route: `/bsky/:did/:postId` exists to handle post/thread retrieval. It will make a request to bsky API, retrieve the data and return a JSON response.

## Prerequisites

- Go (1.14 or higher).
- Environment variables `BSKY_ID` and `BSKY_PW` for authentication or corresponding command-line arguments.
- Package dependencies including `github.com/alexflint/go-arg` and `github.com/joho/godotenv/autoload`.

## Installation

Clone the repository to your local machine and navigate to the project directory:

```bash
git clone https://github.com/hrbrmstr/bskythread.git
cd bskythread
go mod download
go build
# setup .env file with BSKY_ID and BSKY_PW environment variables
./bskythread
```

## Usage

```bash
Usage: bskythread [--identifier BSKY_ID] [--password BSKY_PW] [--listen-port PORT]

Options:
  --identifier BSKY_ID, -i BSKY_ID
                         bsky.social user id [env: BSKY_ID]
  --password BSKY_PW, -p BSKY_PW
                         bsky.social app password [env: BSKY_PW]
  --listen-port PORT, -l PORT
                         port to listen on [default: 3000, env: PORT]
  --help, -h             display this help and exit
```

For this: `https://bsky.app/profile/jayrosen.bsky.social/post/3k3ewlgjhpy2m` post, we can retrieve the JSON via:

```bash
curl -s "http://localhost:3000/bsky/did:plc:3t37x6vfigdzzp2gjcfnzlz4/3k3ewlgjhpy2m"
```

```json
{
  "data": {
    "thread": {
      "$type": "app.bsky.feed.defs#threadViewPost",
      "post": {
        "author": {
          "avatar": "https://cdn.bsky.social/imgproxy/MXa8PEKuW-ukrCKIfk67XhdPSUcZpkMO1iOp7iCiRrs/rs:fill:1000:1000:1:0/plain/bafkreig7vkfbd5c2tzrig5eiuhakch25vm3encxb6j4vrco7vfeommqdxq@jpeg",
          "did": "did:plc:3t37x6vfigdzzp2gjcfnzlz4",
          "displayName": "Jay Rosen",
          "handle": "jayrosen.bsky.social",
          "labels": [],
          "viewer": {
            "blockedBy": false,
            "muted": false
          }
        },
        "cid": "bafyreihgwlirwxye6lrl4mjw4ijvjats24lobruw35blspldu2s2uyktta",
        "embed": {
          "$type": "app.bsky.embed.external#view",
          "external": {
            "description": "G’bye, NATO. G’bye, “Don’t Ask, Don’t Tell” repeal. G’bye, democracy. He’s telling us plain as day. Why aren’t people listening?",
            "thumb": "https://cdn.bsky.social/imgproxy/Qq-rWHvn_KS_Cprfkzx51_DOiGQJKqTl2Cvn0hwCz3A/rs:fit:1000:1000:1:0/plain/bafkreicz7vue54ww5sekfxouvngbixsbhk3vdbiahgj5tkf44jqaq2l5ge@jpeg",
            "title": "People Aren’t Facing Up to the Horrors a New Trump Term Would Bring",
            "uri": "https://newrepublic.com/article/174535/people-arent-facing-horrors-new-trump-term-bring"
          }
        },
        "indexedAt": "2023-07-25T22:20:17.570Z",
        "labels": [],
        "likeCount": 86,
        "record": {
          "$type": "app.bsky.feed.post",
          "createdAt": "2023-07-25T22:20:17.534Z",
          "embed": {
            "$type": "app.bsky.embed.external",
            "external": {
              "description": "G’bye, NATO. G’bye, “Don’t Ask, Don’t Tell” repeal. G’bye, democracy. He’s telling us plain as day. Why aren’t people listening?",
              "thumb": {
                "$type": "blob",
                "mimeType": "image/jpeg",
                "ref": {
                  "$link": "bafkreicz7vue54ww5sekfxouvngbixsbhk3vdbiahgj5tkf44jqaq2l5ge"
                },
                "size": 947415
              },
              "title": "People Aren’t Facing Up to the Horrors a New Trump Term Would Bring",
              "uri": "https://newrepublic.com/article/174535/people-arent-facing-horrors-new-trump-term-bring"
            }
          },
          "langs": [
            "en"
          ],
          "text": "\"Not the odds, but the stakes.\"\n\nThat's my shorthand for the organizing principle we most need in campaign coverage. Not who has what chances of winning, but the consequences for our democracy— given what's possible in 2024. Not the odds, but the stakes.\n\nHere is an example of stakes coverage."
        },
        "replyCount": 1,
        "repostCount": 44,
        "uri": "at://did:plc:3t37x6vfigdzzp2gjcfnzlz4/app.bsky.feed.post/3k3ewlgjhpy2m",
        "viewer": {}
      },
      "replies": [
        {
          "$type": "app.bsky.feed.defs#threadViewPost",
          "post": {
            "author": {
              "avatar": "https://cdn.bsky.social/imgproxy/SUjlP6CnrIQg3BM2StszfY7DZrpfwTNqt-FMR9H4xMc/rs:fill:1000:1000:1:0/plain/bafkreiblpyhlzodiv5fx2pqnyyn7fkyr4ml6f3t3vcv2bsdmq74nee5ocu@jpeg",
              "did": "did:plc:iq2q5xvuoo465bpc6vo4n6u3",
              "displayName": "KansasWoman",
              "handle": "kansaswoman.bsky.social",
              "labels": [],
              "viewer": {
                "blockedBy": false,
                "muted": false
              }
            },
            "cid": "bafyreifkpp2zesk2jlyuyerwd272xlyn6tddwo76olm7jmzpvuhwab2mea",
            "indexedAt": "2023-07-27T02:47:38.330Z",
            "labels": [],
            "likeCount": 0,
            "record": {
              "$type": "app.bsky.feed.post",
              "createdAt": "2023-07-27T02:47:37.900Z",
              "langs": [
                "en"
              ],
              "reply": {
                "parent": {
                  "cid": "bafyreihgwlirwxye6lrl4mjw4ijvjats24lobruw35blspldu2s2uyktta",
                  "uri": "at://did:plc:3t37x6vfigdzzp2gjcfnzlz4/app.bsky.feed.post/3k3ewlgjhpy2m"
                },
                "root": {
                  "cid": "bafyreihgwlirwxye6lrl4mjw4ijvjats24lobruw35blspldu2s2uyktta",
                  "uri": "at://did:plc:3t37x6vfigdzzp2gjcfnzlz4/app.bsky.feed.post/3k3ewlgjhpy2m"
                }
              },
              "text": "THIS THIS THIS"
            },
            "replyCount": 0,
            "repostCount": 0,
            "uri": "at://did:plc:iq2q5xvuoo465bpc6vo4n6u3/app.bsky.feed.post/3k3hvyfmkk42w",
            "viewer": {}
          },
          "replies": []
        }
      ]
    }
  },
  "message": "Success"
}
```

You can test it out on a live server:

```bash
curl -s "http://api.hrbrmstr.de/bsky/did:plc:3t37x6vfigdzzp2gjcfnzlz4/3k3ewlgjhpy2m"
```