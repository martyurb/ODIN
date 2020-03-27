import requests
import bs4

f = open("/home/odin/stats.txt", "a+")
html = requests.get("http://dcufm.redbrick.dcu.ie").text
data = bs4.BeautifulSoup(html, "lxml")
table_row = data.find('body').find_all("tr")[6]
listeners = table_row.find_all("td")[1].text

f.write(__import__("datetime").datetime.now().strftime('%Y-%m-%d (%H:%M:%S) - ') + listeners + "\n")
f.close()
