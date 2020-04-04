# Shomash
A parallel shodan client that can utilize as many shodan API keys as you have (to bypass the 1/request a second rate limit).

If you want to very quickly obtain a shodan passive port scan of an entire enviornment it's super easy.


A basic request:
```
cat ips.txt | shomash > shodan.json
```

To a more advanced combination:
```
subfinder -d "spotify.com"  | zdns A | jq -r 'select(.data.answers[0].type == "A") | .data.answers[].answer' | shomash
```

This single command will pull subdomains from subfinder, check that they resolve, and then pipe them to shomash which will return JSON for each IP passed via stdin.

# Installation
```
go get -u github.com/pry0cc/shomash
```

Then add all your Shodan keys (the more the better!) in `~/.shomash`, Shomash will initialize a new thread for each key, so adding 25 keys will make your scan be able to operate at 25 requests / second, not bad huh?

### Examples

```
echo "1.1.1.1" | shomash
>> {"region_code": null, "ip": 16843009, "postal_code": null, "country_code": "AU", "city": null, "dma_code": null, "last_update": "2020-04-03T23:44:43.219809", "latitude": -33.494, "tags": [], "area_code": null, "country_name": "Australia", "hostnames": ["one.one.one.one"], "org": "Mountain View Communications", "data": [{"_shodan": {"id": "6f9b9363-593c-43aa-9872-91a940234e89", "options": {}, "ptr": true, "module": "dns-udp", "crawler": "82488cbcb7dd25da13f728d04775390417d9ee4e"}, "hash": 1592421393, "os": null, "opts": {"raw": "34ef818500010000000000000776657273696f6e0462696e640000100003"}, "ip": 16843009, "isp": "APNIC and Cloudflare DNS Resolver project", "port": 53, "hostnames": ["one.one.one.one"], "location": {"city": null, "region_code": null, "area_code": null, "longitude": 143.2104, "country_code3": "AUS", "country_name": "Australia", "postal_code": null, "dma_code": null, "country_code": "AU", "latitude": -33.494}, "dns": {"resolver_hostname": null, "recursive": true, "resolver_id": "AMS", "software": null}, "timestamp": "2020-04-03T23:44:43.219809", "domains": ["one.one"], "org": "Mountain View Communications", "data": "\nRecursion: enabled\nResolver ID: AMS", "asn": "AS13335", "transport": "udp", "ip_str": "1.1.1.1"}, {"_shodan": {"id": "2fbc5c4c-909e-4e94-b06c-46781ffa2819", "options": {"hostname": "dxj11.com"}, "ptr": true, "module": "http", "crawler": "122dd688b363c3b45b0e7582622da1e725444808"}, "hash": -1263747863, "os": null, "opts": {}, "ip": 16843009, "isp": "APNIC and Cloudflare DNS Resolver project", "http": {"html_hash": 1937736931, "robots_hash": 1486140972, "redirects": [], "securitytxt": null, "title": "", "sitemap_hash": null, "waf": "CloudFlare", "robots": "User-agent: *\nDisallow:", "favicon": null, "host": "dxj11.com", "html": "<!DOCTYPE html><html data-adblockkey=\"MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANDrp2lz7AOmADaN8tA50LsWcjLFyQFcb/P2Txc58oYOeILb3vBw7J6f4pamkAQVSQuqYsKx3YzdUHCvbVZvFUsCAwEAAQ==_zekC3Aj+GraQRIV8iGLmGfRqKh9SUQQzd9jiiQAaxJkyS9QVuGlvO8pDGPJ/xBTgbjuHJzBtnm1VIpBEk91WEg==\"><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><title></title><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"description\" content=\"See related links to what you are looking for.\"/></head><!--[if IE 6 ]><body class=\"ie6\"><![endif]--><!--[if IE 7 ]><body class=\"ie7\"><![endif]--><!--[if IE 8 ]><body class=\"ie8\"><![endif]--><!--[if IE 9 ]><body class=\"ie9\"><![endif]--><!--[if (gt IE 9)|!(IE)]> --><body><!--<![endif]--><script type=\"text/javascript\">g_pb=(function(){var\nDT=document,azx=location,DD=DT.createElement('script'),aAC=false,LU;DD.defer=true;DD.async=true;DD.src=\"//www.google.com/adsense/domains/caf.js\";DD.onerror=function(){if(azx.search!=='?z'){azx.href='/?z';}};DD.onload=DD.onreadystatechange=function(){if(!aAC&&LU){if(!window['googleNDT_']){}\nLU(google.ads.domains.Caf);}\naAC=true;};DT.body.appendChild(DD);return{azm:function(n$){if(aAC)\nn$(google.ads.domains.Caf);else\nLU=n$;},bq:function(){if(!aAC){DT.body.removeChild(DD);}}};})();g_pd=(function(){var\nazx=window.location,nw={},bH,azw=azx.search.substring(1),aAv,aAw;if(!azw)\nreturn nw;aAv=azw.split(\"&\");for(bH=0;bH<aAv.length;bH++){aAw=aAv[bH].split('=');nw[aAw[0]]=aAw[1]?aAw[1]:\"\";}\nreturn nw;})();g_pc=(function(){var $is_ABP_whitelisted=null;var $Image1=new Image;var $Image2=new Image;var $error1=false;var $error2=false;var $remaining=2;var $random=Math.random()*11;function $imageLoaded(){$remaining--;if($remaining===0)\n$is_ABP_whitelisted=!$error1&&$error2;}\n$Image1.onload=$Image2.onload=$imageLoaded;$Image1.onerror=function(){$error1=true;$imageLoaded();};$Image2.onerror=function(){$error2=true;$imageLoaded();};$Image1.src='/px.gif?ch=1&rn='+$random;$Image2.src='/px.gif?ch=2&rn='+$random;return{azo:function(){return'&abp='+($is_ABP_whitelisted?'1':'0');},$isWhitelisted:function(){return $is_ABP_whitelisted;},$onReady:function($callback){function $poll(){if($is_ABP_whitelisted===null)\nsetTimeout($poll,100);else $callback();}\n$poll();}}})();(function(){var aAo=screen,Rr=window,azx=Rr.location,aAB=top.location,DT=document,Sf=DT.body||DT.getElementsByTagName('body')[0],aAy=0,aAx=0,aAz=0,$IE=null;if(Sf.className==='ie6')\n$IE=6;else if(Sf.className==='ie7')\n$IE=7;else if(Sf.className==='ie8')\n$IE=8;else if(Sf.className==='ie9')\n$IE=9;function aAu($callback){aAz++;aAy=Rr.innerWidth||DT.documentElement.clientWidth||Sf.clientWidth;aAx=Rr.innerHeight||DT.documentElement.clientHeight||Sf.clientHeight;if(aAy>0||aAz>=5){$callback();}\nelse{setTimeout(aAu,100);}}\nvar $num_requirements=2;function $requirementMet(){$num_requirements--;if($num_requirements===0)\naAA();}\naAu($requirementMet);g_pc.$onReady($requirementMet);function aAA(){var ef=undefined,IQ=encodeURIComponent,aAt;if(aAB!=azx&&g_pd.r_s===ef)\naAB.href=azx.href;aAt=DT.createElement('script');aAt.type='text/javascript';aAt.src='/glp'+'?r='+(g_pd.r!==ef?g_pd.r:(DT.referrer?IQ(DT.referrer.substr(0,255)):''))+\n(g_pd.r_u?'&u='+g_pd.r_u:'&u='+IQ(azx.href.split('?')[0]))+\n(g_pd.gc?'&gc='+g_pd.gc:'')+\n(g_pd.cid?'&cid='+g_pd.cid:'')+\n(g_pd.query?'&sq='+g_pd.query:'')+\n(g_pd.search?'&ss=1':'')+\n(g_pd.a!==ef?'&a':'')+\n(g_pd.z!==ef?'&z':'')+\n(g_pd.z_ds!==ef?'&z_ds':'')+\n(g_pd.r_s!==ef?'&r_s='+g_pd.r_s:'')+\n(g_pd.r_d!==ef?'&r_d='+g_pd.r_d:'')+'&rw='+aAo.width+'&rh='+aAo.height+\n(g_pd.r_ww!==ef?'&ww='+g_pd.r_ww:'&ww='+aAy)+\n(g_pd.r_wh!==ef?'&wh='+g_pd.r_wh:'&wh='+aAx)+\n(g_pc.$isWhitelisted()?'&abp=1':'')+\n($IE!==null?'&ie='+$IE:'')+\n(g_pd.partner!==ef?'&partner='+g_pd.partner:'')+\n(g_pd.subid1!==ef?'&subid1='+g_pd.subid1:'')+\n(g_pd.subid2!==ef?'&subid2='+g_pd.subid2:'')+\n(g_pd.subid3!==ef?'&subid3='+g_pd.subid3:'')+\n(g_pd.subid4!==ef?'&subid4='+g_pd.subid4:'')+\n(g_pd.subid5!==ef?'&subid5='+g_pd.subid5:'');Sf.appendChild(aAt);}})();</script></body></html>", "location": "/", "components": {}, "server": "cloudflare", "sitemap": null, "securitytxt_hash": null}, "port": 80, "hostnames": ["one.one.one.one"], "location": {"city": null, "region_code": null, "area_code": null, "longitude": 143.2104, "country_code3": "AUS", "country_name": "Australia", "postal_code": null, "dma_code": null, "country_code": "AU", "latitude": -33.494}, "timestamp": "2020-04-03T23:15:43.261979", "domains": ["one.one"], "org": "Mountain View Communications", "data": "HTTP/1.1 200 OK\r\nDate: Fri, 03 Apr 2020 23:15:43 GMT\r\nContent-Type: text/html; charset=UTF-8\r\nTransfer-Encoding: chunked\r\nConnection: keep-alive\r\nSet-Cookie: __cfduid=dbfe655edc6b348c598416a578cde58581585955742; expires=Sun, 03-May-20 23:15:42 GMT; path=/; domain=.dxj11.com; HttpOnly; SameSite=Lax\r\nX-Adblock-Key: MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBANDrp2lz7AOmADaN8tA50LsWcjLFyQFcb/P2Txc58oYOeILb3vBw7J6f4pamkAQVSQuqYsKx3YzdUHCvbVZvFUsCAwEAAQ==_zekC3Aj+GraQRIV8iGLmGfRqKh9SUQQzd9jiiQAaxJkyS9QVuGlvO8pDGPJ/xBTgbjuHJzBtnm1VIpBEk91WEg==\r\nCF-Cache-Status: DYNAMIC\r\nServer: cloudflare\r\nCF-RAY: 57e67e3ee96a801a-SAN\r\n\r\n", "asn": "AS13335", "transport": "tcp", "ip_str": "1.1.1.1"}], "asn": "AS13335", "isp": "APNIC and Cloudflare DNS Resolver project", "longitude": 143.2104, "country_code3": "AUS", "domains": ["one.one"], "ip_str": "1.1.1.1", "os": null, "ports": [80, 53]}
```

You can do really cool stuff with this such as pulling out RDP images in realtime with jq, identifying all hosts that have RDP exposed, or any other ports from Shodan. 

##### Get subdomains, check they resolve, filter the A records that resolve, pipe IP to shodan, pull shodan data on the first, then return all the IP's that have port 443 open.

```
crobat-client -s spotify.com | zdns A | jq -r 'select(.data.answers[0].type == "A") | .data.answers[].answer' | shomash | jq -r 'select(.ports[] | contains(443)) | .ip_str'
```

```
subfinder -d evernote.com | zdns A | jq -r 'select(.data.answers[0].type == "A") | .data.answers[].answer' | shomash | jq -r 'select(.ports[] | contains(3389)) | .ip_str'
```

You can modify these further to pipe RDP images directly into a file if the device doesn't use NLA for example.


