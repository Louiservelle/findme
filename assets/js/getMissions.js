fetch('http://localhost:80/api/Missions')
    .then(res =>{
        return res.json();
    })
    .then(data =>{
    console.log(data)
    const newH1 = document.createElement("h1");
  
    const newContent = document.createTextNode(data.description);

    newH1.appendChild(newContent);
  
    const currentDiv = document.getElementById("missionsCard");
    document.body.insertBefore(newH1, currentDiv);
});    

