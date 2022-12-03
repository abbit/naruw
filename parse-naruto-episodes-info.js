// this script is used in browser to parse the naruto episodes from https://www.animefillerlist.com/shows/naruto
let tableRows = Array.from(document.querySelector("tbody").children);

let episodes = [];
tableRows.forEach((row) => {
  const cells = row.querySelectorAll("td");
  const episodeNum = cells[0].textContent.trim();
  const title = cells[1].textContent.trim();
  const type = cells[2].textContent.trim();
  episodes.push({
    number: parseInt(episodeNum),
    title: title,
    type: type,
  });
});

let json = JSON.stringify(episodes);
