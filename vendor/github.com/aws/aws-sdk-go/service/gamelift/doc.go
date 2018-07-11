// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package gamelift provides the client and types for making API
// requests to Amazon GameLift.
//
// Amazon GameLift is a managed service for developers who need a scalable,
// dedicated server solution for their multiplayer games. Use Amazon GameLift
// for these tasks: (1) set up computing resources and deploy your game servers,
// (2) run game sessions and get players into games, (3) automatically scale
// your resources to meet player demand and manage costs, and (4) track in-depth
// metrics on game server performance and player usage.
//
// The Amazon GameLift service API includes two important function sets:
//
//    * Manage game sessions and player access -- Retrieve information on available
//    game sessions; create new game sessions; send player requests to join
//    a game session.
//
//    * Configure and manage game server resources -- Manage builds, fleets,
//    queues, and aliases; set auto-scaling policies; retrieve logs and metrics.
//
// This reference guide describes the low-level service API for Amazon GameLift.
// You can use the API functionality with these tools:
//
//    * The Amazon Web Services software development kit (AWS SDK (http://aws.amazon.com/tools/#sdk))
//    is available in multiple languages (http://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-supported.html#gamelift-supported-clients)
//    including C++ and C#. Use the SDK to access the API programmatically from
//    an application, such as a game client.
//
//    * The AWS command-line interface (http://aws.amazon.com/cli/) (CLI) tool
//    is primarily useful for handling administrative actions, such as setting
//    up and managing Amazon GameLift settings and resources. You can use the
//    AWS CLI to manage all of your AWS services.
//
//    * The AWS Management Console (https://console.aws.amazon.com/gamelift/home)
//    for Amazon GameLift provides a web interface to manage your Amazon GameLift
//    settings and resources. The console includes a dashboard for tracking
//    key resources, including builds and fleets, and displays usage and performance
//    metrics for your games as customizable graphs.
//
//    * Amazon GameLift Local is a tool for testing your game's integration
//    with Amazon GameLift before deploying it on the service. This tools supports
//    a subset of key API actions, which can be called from either the AWS CLI
//    or programmatically. See Testing an Integration (http://docs.aws.amazon.com/gamelift/latest/developerguide/integration-testing-local.html).
//
// Learn more
//
//    *  Developer Guide (http://docs.aws.amazon.com/gamelift/latest/developerguide/)
//    -- Read about Amazon GameLift features and how to use them.
//
//    * Tutorials (https://gamedev.amazon.com/forums/tutorials) -- Get started
//    fast with walkthroughs and sample projects.
//
//    * GameDev Blog (http://aws.amazon.com/blogs/gamedev/) -- Stay up to date
//    with new features and techniques.
//
//    * GameDev Forums (https://gamedev.amazon.com/forums/spaces/123/gamelift-discussion.html)
//    -- Connect with the GameDev community.
//
//    * Release notes (http://aws.amazon.com/releasenotes/Amazon-GameLift/)
//    and document history (http://docs.aws.amazon.com/gamelift/latest/developerguide/doc-history.html)
//    -- Stay current with updates to the Amazon GameLift service, SDKs, and
//    documentation.
//
// API SUMMARY
//
// This list offers a functional overview of the Amazon GameLift service API.
//
// Managing Games and Players
//
// Use these actions to start new game sessions, find existing game sessions,
// track game session status and other information, and enable player access
// to game sessions.
//
//    * Discover existing game sessions
//
// SearchGameSessions -- Retrieve all available game sessions or search for
//    game sessions that match a set of criteria.
//
//    * Start new game sessions
//
// Start new games with Queues to find the best available hosting resources
//    across multiple regions, minimize player latency, and balance game session
//    activity for efficiency and cost effectiveness.
//
// StartGameSessionPlacement -- Request a new game session placement and add
//    one or more players to it.
//
// DescribeGameSessionPlacement -- Get details on a placement request, including
//    status.
//
// StopGameSessionPlacement -- Cancel a placement request.
//
// CreateGameSession -- Start a new game session on a specific fleet. Available
//    in Amazon GameLift Local.
//
//    * Match players to game sessions with FlexMatch matchmaking
//
// StartMatchmaking -- Request matchmaking for one players or a group who want
//    to play together.
//
// StartMatchBackfill - Request additional player matches to fill empty slots
//    in an existing game session.
//
// DescribeMatchmaking -- Get details on a matchmaking request, including status.
//
// AcceptMatch -- Register that a player accepts a proposed match, for matches
//    that require player acceptance.
//
// StopMatchmaking -- Cancel a matchmaking request.
//
//    * Manage game session data
//
// DescribeGameSessions -- Retrieve metadata for one or more game sessions,
//    including length of time active and current player count. Available in
//    Amazon GameLift Local.
//
// DescribeGameSessionDetails -- Retrieve metadata and the game session protection
//    setting for one or more game sessions.
//
// UpdateGameSession -- Change game session settings, such as maximum player
//    count and join policy.
//
// GetGameSessionLogUrl -- Get the location of saved logs for a game session.
//
//    * Manage player sessions
//
// CreatePlayerSession -- Send a request for a player to join a game session.
//    Available in Amazon GameLift Local.
//
// CreatePlayerSessions -- Send a request for multiple players to join a game
//    session. Available in Amazon GameLift Local.
//
// DescribePlayerSessions -- Get details on player activity, including status,
//    playing time, and player data. Available in Amazon GameLift Local.
//
// Setting Up and Managing Game Servers
//
// When setting up Amazon GameLift resources for your game, you first create
// a game build (http://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// and upload it to Amazon GameLift. You can then use these actions to configure
// and manage a fleet of resources to run your game servers, scale capacity
// to meet player demand, access performance and utilization metrics, and more.
//
//    * Manage game builds
//
// CreateBuild -- Create a new build using files stored in an Amazon S3 bucket.
//    To create a build and upload files from a local path, use the AWS CLI
//    command upload-build.
//
// ListBuilds -- Get a list of all builds uploaded to a Amazon GameLift region.
//
// DescribeBuild -- Retrieve information associated with a build.
//
// UpdateBuild -- Change build metadata, including build name and version.
//
// DeleteBuild -- Remove a build from Amazon GameLift.
//
//    * Manage fleets
//
// CreateFleet -- Configure and activate a new fleet to run a build's game servers.
//
// ListFleets -- Get a list of all fleet IDs in a Amazon GameLift region (all
//    statuses).
//
// DeleteFleet -- Terminate a fleet that is no longer running game servers or
//    hosting players.
//
// View / update fleet configurations.
//
// DescribeFleetAttributes / UpdateFleetAttributes -- View or change a fleet's
//    metadata and settings for game session protection and resource creation
//    limits.
//
// DescribeFleetPortSettings / UpdateFleetPortSettings -- View or change the
//    inbound permissions (IP address and port setting ranges) allowed for a
//    fleet.
//
// DescribeRuntimeConfiguration / UpdateRuntimeConfiguration -- View or change
//    what server processes (and how many) to run on each instance in a fleet.
//
//    * Control fleet capacity
//
// DescribeEC2InstanceLimits -- Retrieve maximum number of instances allowed
//    for the current AWS account and the current usage level.
//
// DescribeFleetCapacity / UpdateFleetCapacity -- Retrieve the capacity settings
//    and the current number of instances in a fleet; adjust fleet capacity
//    settings to scale up or down.
//
// Autoscale -- Manage auto-scaling rules and apply them to a fleet.
//
// PutScalingPolicy -- Create a new auto-scaling policy, or update an existing
//    one.
//
// DescribeScalingPolicies -- Retrieve an existing auto-scaling policy.
//
// DeleteScalingPolicy -- Delete an auto-scaling policy and stop it from affecting
//    a fleet's capacity.
//
// StartFleetActions -- Restart a fleet's auto-scaling policies.
//
// StopFleetActions -- Suspend a fleet's auto-scaling policies.
//
//    * Manage VPC peering connections for fleets
//
// CreateVpcPeeringAuthorization -- Authorize a peering connection to one of
//    your VPCs.
//
// DescribeVpcPeeringAuthorizations -- Retrieve valid peering connection authorizations.
//
//
// DeleteVpcPeeringAuthorization -- Delete a peering connection authorization.
//
// CreateVpcPeeringConnection -- Establish a peering connection between the
//    VPC for a Amazon GameLift fleet and one of your VPCs.
//
// DescribeVpcPeeringConnections -- Retrieve information on active or pending
//    VPC peering connections with a Amazon GameLift fleet.
//
// DeleteVpcPeeringConnection -- Delete a VPC peering connection with a Amazon
//    GameLift fleet.
//
//    * Access fleet activity statistics
//
// DescribeFleetUtilization -- Get current data on the number of server processes,
//    game sessions, and players currently active on a fleet.
//
// DescribeFleetEvents -- Get a fleet's logged events for a specified time span.
//
// DescribeGameSessions -- Retrieve metadata associated with one or more game
//    sessions, including length of time active and current player count.
//
//    * Remotely access an instance
//
// DescribeInstances -- Get information on each instance in a fleet, including
//    instance ID, IP address, and status.
//
// GetInstanceAccess -- Request access credentials needed to remotely connect
//    to a specified instance in a fleet.
//
//    * Manage fleet aliases
//
// CreateAlias -- Define a new alias and optionally assign it to a fleet.
//
// ListAliases -- Get all fleet aliases defined in a Amazon GameLift region.
//
// DescribeAlias -- Retrieve information on an existing alias.
//
// UpdateAlias -- Change settings for a alias, such as redirecting it from one
//    fleet to another.
//
// DeleteAlias -- Remove an alias from the region.
//
// ResolveAlias -- Get the fleet ID that a specified alias points to.
//
//    * Manage game session queues
//
// CreateGameSessionQueue -- Create a queue for processing requests for new
//    game sessions.
//
// DescribeGameSessionQueues -- Retrieve game session queues defined in a Amazon
//    GameLift region.
//
// UpdateGameSessionQueue -- Change the configuration of a game session queue.
//
// DeleteGameSessionQueue -- Remove a game session queue from the region.
//
//    * Manage FlexMatch resources
//
// CreateMatchmakingConfiguration -- Create a matchmaking configuration with
//    instructions for building a player group and placing in a new game session.
//
//
// DescribeMatchmakingConfigurations -- Retrieve matchmaking configurations
//    defined a Amazon GameLift region.
//
// UpdateMatchmakingConfiguration -- Change settings for matchmaking configuration.
//    queue.
//
// DeleteMatchmakingConfiguration -- Remove a matchmaking configuration from
//    the region.
//
// CreateMatchmakingRuleSet -- Create a set of rules to use when searching for
//    player matches.
//
// DescribeMatchmakingRuleSets -- Retrieve matchmaking rule sets defined in
//    a Amazon GameLift region.
//
// ValidateMatchmakingRuleSet -- Verify syntax for a set of matchmaking rules.
//
// See https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01 for more information on this service.
//
// See gamelift package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/gamelift/
//
// Using the Client
//
// To contact Amazon GameLift with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon GameLift client GameLift for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/gamelift/#New
package gamelift
