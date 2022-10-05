package bqproto

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/bigquery/storage/managedwriter/adapt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func SchemaDescriptorProto(project, dataset, table string) (bigquery.Schema, *descriptorpb.DescriptorProto, error) {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, project)
	if err != nil {
		return nil, nil, err
	}

	defer client.Close()

	meta, err := client.Dataset(dataset).Table(table).Metadata(ctx)
	if err != nil {
		return nil, nil, err
	}

	// for _, f := range meta.Schema {
	// 	if f.Type == bigquery.TimeFieldType {
	// 		f.Type = bigquery.StringFieldType
	// 	}
	// }

	tschema, err := adapt.BQSchemaToStorageTableSchema(meta.Schema)
	if err != nil {
		return nil, nil, err
	}

	pdesc, err := adapt.StorageSchemaToProto2Descriptor(tschema, table)
	if err != nil {
		log.Printf("%s: %v", meta.Name, tschema)
		return nil, nil, err
	}

	mdesc, ok := pdesc.(protoreflect.MessageDescriptor)
	if !ok {
		return nil, nil, err
	}

	desc, err := adapt.NormalizeDescriptor(mdesc)
	if err != nil {
		return nil, nil, err
	}

	return meta.Schema, desc, nil
}
