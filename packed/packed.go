package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/4SVeVBTWRbGHxiUCVtE1maURpYQGILIpoDsS9sQTAGCAqZBZNMAESOromAEIQkBFVDZgkbaAIKBCEIMRLYxgWGLsoUomMjiBLENBBRoppgeW5yemj5/vFe37r1f1T31O9+H9NwGUgFkAVmAehXnC2wpdeAvQFhcbER0pOlvP5OwOEwyHBsXgz7qJwNIaRdXh+JuIOJ4ByCZC2Jjeq8Ow9H5cMJJ2+0EfkgQWSbYIia+qmKbfq7e3jXkytTopyUZGNQqkH5Io6yIEoGyN7frsbezmwhfjmqKChUMXZJpKgnCwoPCTzUbqu2SFmTzFB+Tqoeb7V+AhisZYa0kz+YAFziiiJT66RaexSBn1hkk0ctH8x7WFx7OAiLHxrG7A8FatzPAafnZrO+aDAWKN66zmDMGruqvtuMije3GN4gWEmbE484RhdFjwvajLFq4kwFPzgY2g0UNsIUpB4v74uA5eKik/lz/Bhj+i5OhxbO6CDcnwIpoRpI+IWPKnHVp5PT1V6j5dL5+ENJxO9Ly3LFn0yaVWTlHZZhndU3Pa46PvUkJ8IIG8F8VJp3Xlz+auqxHfZu7HyZxss61ct+pQUCUI37uIjhioo5UKKHRiE+dqF8/x68SL0eASdAkBGshi37dCDJ5LAv+zrUcaErVwjZywEnug0vUQM06/mKvjZfNSoKTI43vUfboev0PH8WXP88/Pcl4HDWIH3xGxwhBcusucWLh+8T1q4gRKx/XVhh592nObapRnrNDK//sQYPg3Pm1fuuO129PecsFDsRrHvz1hvYcdhh3AI9jhdBUGWaSs2kTGtpVL51+QK0Dt4AGRItiQ85lni3J6EcYdnQnMtPjPtFdGhGpPJNuyY4JwxrbSAfRvW+B4oPcy0FeH1xaDlYkm1hITST33G1d9H+Q/sm+3Y1dok7smH1zIIy53LG/nxES3C//cpZ0BWuezUzZTXHJEOgRB07Gdt1BRdFX3+lVZhUtBs4tau4+07m86MtZkAKAjQ2k5w5ZSc9ptUhpAAgFAcBXaJdTv4VW7b+hhZ8KT/jKrCC7OvSCDyKOt0+eKQ6QfVehhRz1rr8ffTnEuFFKKy+0607x1BX0YCNyzGdpUG996QqFfpjWuKODRHHNX2CafzzX0KzTFLVnwDFa0JZxpk8h0NMfB1fvOumsU/rkPQarc6Zm/iIi9GGL1vTnPA6ffI7aYvfy4wt8e6LDbBvH1qLttoHx0P7onZOEmqMjFWzyzZ1+H9yh4DMXHDWNvH4EN+Uaa1diunXRRExUK4zaS8p/qVwORn4nbPdn0cJR3TQot6fmeQ66vqQfXgfPwbOxkjnxGn5S2fqk41v+aZSKChl9s2VXlHGSyG2BnVNS7mJCkf0phyzspd05XkO+0D021VyVgRorqvMZ4HS+QPv3wpLw0zcR/1CeWRadjWxScbeapPWdAcfWhUWM8lIPp4fNMyBZRBimwTpNz9iHpeY28vNIhz2in9ClkToCN5RHL+M96mcp16a0iQLxp/T54btu0CGgoX3oxsUG9Zz8gJdKAn3+K4gipzaGIjvhlnf+QEQJtUoa1s14YfW+2etQUXZhXe2gdq7zdgOHaY9Mm3tPcxNhUVcejuqLMqzR7ci7p4azu/eaIy75VSV6m4rXVJ/VVqYXQp8pjmwEyHPSLZkag6LqXSW7a1ZydTN0SOvXTefagC/ACF3DVIVSAFC+bSswi5nfAqP8B2B+h4VQWh2K8z0Sx7OXZ4p9NUQVbdoDbO8My6+0jG/SYtaIPB6/BEV8ntvBjlauGU3Y1+sBFgq4XDfJ0tQUnAFl6Mvd99wAx159JAxu5GCaDdV8cb90xniO/73mfK2cZfs4PXEi0axgNe15QbSYVkhiQuyz0laiM6CrpUpjRjzMW9nJvUHFjfo9ZD/Yg0mlUsgwXV/VOBiGSzX1uva3pwqgfv+U+nLCfgl++vHdEembhTisquz4B7LyvEyx9zQa9Whucf1jilVbs1Ft2UpP36FLec6DV0RsUFsLHAwy1Arbo98MRtmMXUUXnEA5mxQQXl8jq0yI79hWke26RftoVWRzuEKZ10xNiEzx9/LdlS/Q/sG6IA9J8EA6qZ5UA7FYmPNTt8x4rQTjLvBvOEP8sqzePLHIePQf+zIYSxAsPjnCPftPw6lJazlc2UJrNkXdOwd0SZ5CmkltDepRSD6ofEtVsV8SvvGEbpe/yutyGov9Tv8eBhZiEm3whGqE9HHYExCcq1B0b6/2RfO1RL9zmm5rDg7Z3MBdd9mQ0Ef0fIaZZFAjlqDEffd9wYM0rno6z+BhRwxpUtLnq11x3UYESdHdO+S+y5l7hyJmFXc2mqXoBuFMjkMRyvG22XY7vT64jP/1OMqT6Ggbl9y58LRgnXVRvUu/q7SIeO/5m1Nlfauht2di9tUKMqM/QtX5XoITzJbZgcmHRLIomXsYkfBTaZfogysKWtpG4hJWkK/oa6pfKHwzMIQ5JA0ANltsCwCe47+lcPvvFP6bPOr96tDNy1uPID2lpFW2fY3qrX64GdVfipK++f2z4N6qtumhW4dF7Ru1x9LAnznqVrHN+dr6ZuVvxKZBwP+btq9C/7tFvxUE2HBC7wD+2DCZ7ZvbIAAE6EgBAHXH5upfAQAA//9SCSvp4AgAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
