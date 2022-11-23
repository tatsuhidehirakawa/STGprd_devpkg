import React from "react"
import "./SiteTop.scss"
import ApiFetch from '../../hooks/ApiFetch.jsx';

const SiteTop = () => {
  return (
    <div className="sitetop--outline">

      <div className="sitetop--inlineA">
      {/* <p className="sitetop__textA">- <b>SiteTop</b> -</p> */}
      <div className="sitetop--inlineA--boxA__pictureA">
      </div>
      <div className="sitetop--inlineA--boxB__textA">
        <p className="sss">
          It is a full-stack package for web development, 
          focusing on the technology stack presented in the next block. 
          It is optimized primarily as a template for creating large-scale CtoC sites. 
          In particular, it is RESTful,
          does not use CSS-in-JS, minimizes Paas dependencies, 
          emphasizes orthogonality and loose coupling, static typing, 
          and Docker usage.
        </p>
      </div>
      </div>

      <div className="sitetop--inlineB">
        <div className="sitetop--inlineB--boxA__pictureA">
          <p>
            Technology stack:<br />
            Frontend: React.js, TypeScript, Sass<br />
            Backend: Golang(Gin,sqlc), PostgreSQL<br />
            <b>Dev</b>lopment environment: Docker<br />
            <b>T</b>e<b>st</b> Environment: CircleCI, Docker(Multistage Build)<br />
            <b>St</b>a<b>g</b>ing Environment: Terraform
          </p>
        </div>
        <div className="sitetop--inlineB--boxB__textA">
          <div className="sitetop__textS"></div>
          <div className="sitetop__textS"></div>
          <div className="sitetop__textS"></div>
          <ApiFetch />
          <div className="sitetop__textS"></div>
          <div className="sitetop__textS"></div>
          <div className="sitetop__textS"></div>
          <div className="sitetop__textS"></div>
        </div>

      </div>
      <div className="sitetop--inlineC__pictureA">
        </div>
    </div>
  );
}

export default SiteTop;