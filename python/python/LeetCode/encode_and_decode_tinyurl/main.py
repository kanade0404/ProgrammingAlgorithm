import string
import random


class Codec:
    original_url = ""
    short_url = ""
    def encode(self, longUrl: str) -> str:
        """Encodes a URL to a shortened URL.
        """
        self.original_url = longUrl
        self.short_url = "".join([random.choice(string.ascii_letters + string.digits) for i in range(10)])
        print(self.short_url, self.original_url)
        return self.short_url

    def decode(self, shortUrl: str) -> str:
        """Decodes a shortened URL to its original URL.
        """
        print(self.short_url, shortUrl)
        if self.short_url == shortUrl:
            return self.original_url
        return ""

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.decode(codec.encode(url))