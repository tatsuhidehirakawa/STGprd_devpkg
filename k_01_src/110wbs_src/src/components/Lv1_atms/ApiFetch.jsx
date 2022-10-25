import React, {useState, useEffect} from 'react'

const ApiFetch = () => {
    console.log('ApiFetch.jsx: ApiFetch()')   // Debug.
    const [posts, setPosts] = useState([])

    useEffect(() => {
        console.log('ApiFetch.jsx: useEffect()')   // Debug.
        fetch('https://jsonplaceholder.typicode.com/posts', {method: 'GET'})
        // fetch('http://localhost:8080/somePost', {method: 'GET'})
        // fetch('http://localhost:8080/someGet', {method: 'GET'})
        .then(res => res.json())
        .then(data => {
            setPosts(data)
        })
        .catch(err => {
            console.log('ApiFetch.jsx: useEffect(): err: ', err)
        });
    },[])

    return (
        <div>
            <ul>
                {
                    posts.map(post => <li key={post.id}>{post.title}</li>)
                }
            </ul>
            
        </div>
    )
}

export default ApiFetch