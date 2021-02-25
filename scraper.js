const myArgs = process.argv.slice(2);

const puppeteer = require("puppeteer");

(async () => {
  const browser = await puppeteer.launch({
    headless: true
  });
  const page = await browser.newPage();
  // hacky little thing to get over 30 songs, still limited to 101 - TODO
  await page.setViewport({ width: 1366, height: 10000 });
  await page.goto(myArgs[0]);

  await page.waitForTimeout(1000)
  // Could also fetch album artist here - TODO
  let r = await page.$$eval("div[data-testid='tracklist-row'] div[aria-colindex='2'] div div",  els => els.map( el => el.textContent ));
  console.log(r.toLocaleString());
  await browser.close();
})();
