import './app.css'
import 'vite/modulepreload-polyfill'
import 'htmx.org'
import Alpine from 'alpinejs'
import Swal from 'sweetalert2'

window.Alpine = Alpine
Alpine.start()

document.addEventListener('htmx:confirm', (evt) => {
    console.log(evt)
    if (!evt.target.hasAttribute('hx-confirm')) {
        return
    }

    evt.preventDefault()

    Swal.fire({
        title: evt.detail.question
    }).then( (result) => {
        if (result.isConfirmed) {
            evt.detail.issueRequest(true);
        }
    } )
})