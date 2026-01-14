            fetch("http://localhost:8080/api/activities")
                .then(res => res.json())
                .then(data => {
                    const activities = data.Response.activities;
                    const parseActivities = activities.map(a => ({
                        date: new Date(a.period),
                        activity: a.activityDetails.directorActivityHash,
                        kills: a.values.kills.basic.value, 
                        deaths: a.values.deaths.basic.value,
                        completed: a.values.completed.statId,
                        timePlayed: a.values.timePlayedSeconds.basic.displayValue

                    }));
                    console.log(parseActivities);

                    const results = document.getElementById("results")
                    results.innerHTML="";
                    parseActivities.forEach(a => {
                        const div = document.createElement("div");
                        div.textContent = `
                        Date: ${a.date.toLocaleString()}
                        Activity: ${a.activity}
                        Kills: ${a.kills}
                        Deaths: ${a.deaths}
                        Time: ${a.timePlayed}
                        `
                        ;
                    
                        results.appendChild(div);
                    })
                })
                .catch(err => console.log(err));