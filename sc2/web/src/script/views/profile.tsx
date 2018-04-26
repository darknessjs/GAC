
import * as React from 'react';
import * as jQuery from 'jquery';


type Props = {match: any};

type Sc2Profile = {
    displayName: string,
    clanName: string
}

class Profile extends React.Component<Props> {

    public state = {sc2Profile: null}
    public componentDidMount() {
        const name = this.props.match.params.name;
        const sc2Id = this.props.match.params.sc2Id;
        jQuery.get("/action/sc2/profile/" + sc2Id + "/1/" + name + "/", {}, (data) => {
            this.setState({
                sc2Profile: JSON.parse(data),
            })
        })
    }
    private renderProfile() {
        if (this.state.sc2Profile == null) return '';
        const sc2Profile:Sc2Profile = this.state.sc2Profile || {
            displayName: '',
            clanName: '',
        };
        const displayName = sc2Profile.displayName;
        const clanName = sc2Profile.clanName;
        return (
            <div>
                <div>
                    显示名
                </div>
                <div>
                    {displayName}
                </div>
                <div>
                    战队名
                </div>
                <div>
                    {clanName}
                </div>
            </div>
        )
    }
    public render() {
        console.log(this);

        return (
            <div className="result-main">
                登录成功的个人资料界面, {this.props.match.params.name}
                {this.renderProfile()}
            </div>
        );
    }
}


export default Profile;
