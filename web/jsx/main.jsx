var React = require('react');
var ReactDOM = require('react-dom');

var Header = require('./header');
var Characters = require('./characters');
var Error = require('./error');
var Message = require('./message');

var Main = React.createClass({
    getInitialState: function () {
        return {
            characters: [],
            showMessage: false,
            message: '',
            showErrors: false,
            errorMessage: ''
        };
    },

    saveCharacters: function(characters){
        var xhr = new XMLHttpRequest();

        xhr.open('POST', encodeURI('api/characters'));

        xhr.onload = function () {
            if (xhr.status === 200) {
                this.setState({ showErrors: false, showMessage: true, message: "Characters have been updated" })
                return;
            }
            else if (xhr.status === 404) {
                this.setState({ showMessage: false, showErrors: true, errorMessage: "Darn, something went wrong - " + xhr.responseText})
                return;
            } else if (xhr.status !== 200) {
                this.setState({ showMessage: false, showErrors: true, errorMessage: "Darn, something went wrong - " + xhr.responseText })
                return;
            }
        }.bind(this);
        xhr.send(JSON.stringify(characters));
    },

    getCharacters: function () {
        var xhr = new XMLHttpRequest();

        xhr.open('GET', encodeURI('api/characters'));

        xhr.onload = function () {
            if (xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);
                this.setState({showErrors: false, characters: data});
                return;
            }
            else if (xhr.status === 404) {
                this.setState({ showErrors: true, errorMessage: "Darn, something went wrong" })
                return;
            } else if (xhr.status !== 200) {
                this.setState({ showErrors: true, errorMessage: "Darn, something went wrong" })
                return;
            }
        }.bind(this);
        xhr.send();
    },

    componentDidMount: function() {
        this.getCharacters();
    },

    render: function () {

        var error;
        if (this.state.showErrors) {
            error = <Error errorMsg={this.state.errorMessage}/>
        }

        var message;
        if (this.state.showMessage) {
            message = <Message message={this.state.message}/>
        }

        return (
            <div>
                <Header />
                <main>
                    {error}
                    {message}
                    <div>
                       <Characters characters={this.state.characters} saveCharacters={this.saveCharacters} />
                    </div>
                </main>
            </div>
        )
    }
});

var homePage = React.createElement(Main);

ReactDOM.render(homePage, document.querySelector('#wrapper'));
