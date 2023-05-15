
websites = [
    "filecoin.io",
    "libp2p.io",
    "ipld.io",
    "protocol.ai",
    "docs.ipfs.tech",
    "tldr.filecoin.io",
    "green.filecoin.io",
    "drand.love",
    "web3.storage",
    "consensuslab.world",
    "strn.network",
    "research.protocol.ai",
    "ipfs.tech",
    "blog.ipfs.tech",
    "docs.libp2p.io",
    "blog.libp2p.io",
    "en.wikipedia-on-ipfs.org/wiki"
]
types = ["KUBO", "HTTP"]
regions = [
    "eu-central-1",
    "us-west-1",
    "us-east-2",
    "ap-southeast-2",
    "sa-east-1",
    "af-south-1",
    "ap-south-1",
    ]

for website in websites:
    for type in types:
        for region in regions:
            variant = f"""- website: {website}
  type: {type}
  region: {region}"""
            print(variant)