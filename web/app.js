            fetch("http://localhost:8080/api/activities")
                .then(res => res.json())
                .then(data => {
                    console.log(data);
                })
                .catch(err => console.log(err));