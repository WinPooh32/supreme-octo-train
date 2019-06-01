import * as React from "react";
import { Segment, Grid, List, Checkbox, Divider, Icon, Container, Pagination, PaginationProps } from "semantic-ui-react";
import { ItemsList } from "./ItemsList";
import { FilterSearch } from "./FilterSearch";
import { Redirect, matchPath } from "react-router";
import { isNumber } from "util";

const regexNumber = /^\d+$/g;

export class ViewItems extends React.Component<any, any> {
    match: any

    constructor(props: any, state: any) {
        super(props, state)

        // this.match = props.match

        // this.state = { items: props.items }
        this.handlePageChange = this.handlePageChange.bind(this)
    }

    handlePageChange(event: React.MouseEvent<HTMLAnchorElement>, data: PaginationProps) {
        window.location.href = `../page/${data.activePage}`
    }

    render() {
        // let params = new URLSearchParams(location.search);

        let page = this.props.match.params.id
        let pageNum = parseInt(page, 10)

        console.log("ViewItems", page)

        let matchNumber = regexNumber.exec(page)

        if (matchNumber === null || matchNumber.length === 0) {
            return <Redirect to="/page/1" />
        }

        return (
            <div>
                {/* <Segment>
                    <Grid columns={2} relaxed='very'>
                        <Grid.Column verticalAlign="middle">
                            <FilterSearch />
                        </Grid.Column>
                        <Grid.Column>
                            <List>
                                <List.Item><Checkbox label='Новый' defaultChecked /></List.Item>
                                <List.Item><Checkbox label='Заканчивается' defaultChecked /></List.Item>
                                <List.Item><Checkbox label='Нет на складе' defaultChecked /></List.Item>
                            </List>
                        </Grid.Column>
                    </Grid>

                    <Divider vertical>  <Icon name='filter' /> </Divider>
                </Segment> */}

                <ItemsList items={this.props.items} page={pageNum} />

                {/* <Container textAlign="center">
                    <Pagination
                        defaultActivePage={pageNum}
                        onPageChange={this.handlePageChange}
                        ellipsisItem={{ content: <Icon name='ellipsis horizontal' />, icon: true }}
                        firstItem={{ content: <Icon name='angle double left' />, icon: true }}
                        lastItem={{ content: <Icon name='angle double right' />, icon: true }}
                        prevItem={{ content: <Icon name='angle left' />, icon: true }}
                        nextItem={{ content: <Icon name='angle right' />, icon: true }}
                        totalPages={40}
                    />
                </Container> */}
            </div>
        );
    }
}