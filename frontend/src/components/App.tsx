import * as React from "react";
import { Container, Dimmer, Loader } from "semantic-ui-react";
import { MyMenu } from "./MyMenu";
import { Switch, Route, RouteComponentProps, Redirect } from "react-router";
import { ViewItems } from "./ViewItems";
import { ViewItemInfo } from "./ViewItemInfo";


function About(): React.ReactComponentElement<any, any> {
    return <h2>About</h2>;
}

function Users(): React.ReactComponentElement<any, any> {
    return <h2>Users</h2>;
}

export class App extends React.Component<any, any> {
    constructor(props: any, state: any) {
        super(props, state)

        this.state = { items: [], loaded: false }

        fetch("../forecast")
            .then((value) => {
                if (value.status == 200) {
                    value.json().then((v) => {
                        console.log("downloaded:", v)

                        setTimeout(() => {
                            this.setState((state: any, props: any) => ({
                                items: v,
                                loaded: true
                            }));
                        }, 1000)

                    })
                }
            })
            .catch((err) => {
                console.log(err)
            });

    }

    render() {
        if (!this.state.loaded) {
            return (
                <Dimmer active>
                    <Loader />
                </Dimmer>
            )
        }

        return (
            <Container>
                <MyMenu />

                <Switch>
                    <Route path="/page/:id" render={(props) => {
                        return <ViewItems match={props.match} items={this.state.items} />
                    }} />

                    <Route path="/item/:id" render={(props) => {
                        let id = props.match.params.id

                        if (this.state.items === null || id < 1 || id > this.state.items.length) {
                            return <Redirect to="../page/1" />
                        } else {
                            console.log("Route id ", id - 1)
                            return <ViewItemInfo item={this.state.items[id - 1]} />
                        }
                    }} />

                    <Route path="/about/" component={About} />
                    <Route path="/users/" component={Users} />

                    {/* if not matched any route */}
                    <Route render={(props) => {
                        return <Redirect to="../page/1" />
                    }} />

                </Switch>

            </Container>
        )
    }
}