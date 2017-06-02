import bootstrap from 'bootstrap';

import {Component} from 'react';

class Hello extends Component {
    render() {
        return (
            <h1>hello!world!</h1>
        )
    }
}

ReactDOM.render(
    <Hello />,
    document.querySelector('main')
);
