import { useState } from 'react'
import './App.css'
import catalog from './catalog.json'

function imageUrl(id: string): string {
  return `https://static.wixstatic.com/media/${id}/v1/fill/w_339,h_191,q_90/${id}`
}

function App() {
  const [videoId, setVideoId] = useState<string>("")

  if (videoId !== "") {
    return (
      <>
        <iframe 
        width={window.innerWidth} 
        height={window.innerHeight* 0.7} 
        src={`https://www.youtube.com/embed/${videoId}?si=PbZiaLHw2SDOLMIc?modestbranding=1`}
        title="YouTube video player"       
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
        allowFullScreen></iframe>
        <button onClick={() => { setVideoId("") }}>Back</button>
      </>
    )
  }

  return (
    <div className="container">
    <div className="header">
      <h1>
      FERANO WIBISONO
      </h1>
      <h2>
      FILM DIRECTOR & DIRECTOR OF PHOTOGRAPHY
      </h2>
    </div>
    {catalog.map((item) => <img width="339" onClick={() => { setVideoId(item.videoId) }} src={imageUrl(item.customPoster)}/>)}
    </div>
  )
}

export default App
