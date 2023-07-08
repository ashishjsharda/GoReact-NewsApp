import React,{useEffect,useState} from "react";

function App(){
  const [articles,setArticles] = useState([]);
    useEffect(()=>{
        fetch("http://localhost:8080/top-articles")
        .then(res=>res.json())
        .then(articles=>setArticles(articles));
    },[]);
  return (
      <div>
        {articles.map((article, index) => (
            <div key={index}>
              <h2>{article.title}</h2>
              <a href={article.url}>Read more</a>
            </div>
        ))}
      </div>
  );
}

export default App;