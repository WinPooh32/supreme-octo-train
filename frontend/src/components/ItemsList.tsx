import * as React from "react";
import { List, ListItemProps } from 'semantic-ui-react'
import { Link } from "react-router-dom";

export interface ItemsListProps { items: Array<any>; page: number }

export class ItemsList extends React.Component<ItemsListProps, {}> {
    render() {
        let items = new Array<React.ReactNode>()
        let rawlist = this.props.items

        for (let i in rawlist) {
            let item = rawlist[i]

            items.push(
                <List.Item key={item.id} >
                    <List.Content>
                        {/* href={`/?item=${item.id}`} */}
                        <List.Header><Link to={`../item/${item.id + 1}`}>{item.name}</Link></List.Header>
                        {/* <List.Description>512 шт</List.Description> */}
                    </List.Content>
                </List.Item>
            )
        }

        return (
            <List relaxed link={true}>
                {items}
            </List>
        )
    }
}
