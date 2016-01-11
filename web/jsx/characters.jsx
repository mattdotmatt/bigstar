var React = require('react');
var $ = require('jQuery');

module.exports = React.createClass({

    handleSubmit: function(e) {

        e.preventDefault();

        var characters = $('#characters tr:has(td)').map(function(i, v) {
            var $td =  $('td', this);

            return {
                firstName: $td.eq(0)[0].firstChild.value,
                lastName: $td.eq(1)[0].firstChild.value,
            }
        }).get();

        this.props.saveCharacters(characters);

        return;
    },

    render: function () {
        return <form onSubmit={this.handleSubmit}>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>First name</th>
                        <th>Last name</th>
                    </tr>
                </thead>
                <tbody id="characters">
                    {
                        this.props.characters.map(function (c) {
                            return <tr>
                                    <td><input type="text" name="people[][firstname]" defaultValue={c.firstName}/></td>
                                    <td><input type="text" name="people[][surname]" defaultValue={c.lastName}/></td>
                                </tr>
                        })
                    }
                </tbody>
            </table>
            <input type="submit" value="OK" className="btn btn-success pull-right"/>
        </form>
    }
});